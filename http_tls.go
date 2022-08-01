package main

import (
	"log"
	"net/http"
)

// openssl genrsa -out server.key 2048
// openssl req -nodes -new -key server.key -subj "/CN=localhost" -out server.csr
// openssl x509 -req -sha256 -days 365 -in server.csr -signkey server.key -out server.crt

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("hello, world!"))
}

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServeTLS(
		":8000",
		"~/.ssl/server.crt",
		"~/.ssl/server.key", nil); err != nil {
		log.Fatalln(err)
	}
}
