# Go Http File server

A simple HTTP file server written in Go that allows you to:

- Browse files in the current directory
- Download files via clickable links
- Upload files directly from the browser
- Accessible over your local network (`0.0.0.0`)
- Built for a Windows/Linux file transfer binary for pen testing

---

## Requirements

- [Go](https://golang.org/dl/) 1.18 or later
- Windows, macOS, or Linux

```bash
git clone https://github.com/Ilesnat/Win_HTTP_Server
cd Win_HTTP_Server
go build -o http_server main.go
# GOOS=windows GOARCH=amd64 go build -o file-server.exe main.go <- For windows
```
