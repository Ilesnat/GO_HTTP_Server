package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const directory = "./" // current directory

func main() {
	// CXhange this to any interface and port of your choosing
	listener := "0.0.0.0:8080"
	fmt.Printf("Listening on %s\n", listener)

	mux := http.NewServeMux()

	mux.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir(directory))))

	// Homepage: file browser + upload
	mux.HandleFunc("/", upload)

	server := &http.Server{
		Addr:           listener,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseMultipartForm(20 << 20) // 20 MB max
		file, header, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Error reading file: "+err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		dstPath := filepath.Join(directory, header.Filename)
		dst, err := os.Create(dstPath)
		if err != nil {
			fmt.Println(err)
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			fmt.Println(err)
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	files, err := os.ReadDir(directory)
	if err != nil {
		http.Error(w, "Can not list files: ", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, `<html><body>`)
	fmt.Fprintln(w, `<h1>File Browser</h1>`)
	fmt.Fprintln(w, `<form action="/" method="post" enctype="multipart/form-data">
	<input type="file" name="file">
	<input type="submit" value="Upload">
	</form><hr>`)

	fmt.Fprintln(w, "<h2>Files in current directory:</h2><ul>")
	for _, f := range files {
		name := f.Name()
		fmt.Fprintf(w, `<li><a href="/files/%s">%s</a></li>`, name, name)
	}
	fmt.Fprintln(w, "</ul></body></html>")
}
