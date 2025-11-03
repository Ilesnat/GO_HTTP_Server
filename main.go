package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	listner := "0.0.0.0:8080"

	fmt.Printf("Listening on %s\n", listner)

	dirToServe := "./"
	fs := http.FileServer(http.Dir(dirToServe))
	http.Handle("/", fs)

	s := &http.Server{
		Addr:           listner,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
