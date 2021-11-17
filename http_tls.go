package main

import (
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte("hello, world!"))
}

func main() {
    http.HandleFunc("/", handler)
    if err := http.ListenAndServeTLS(":8000", "cert.pem", "key.pem", nil); err != nil {
        log.Fatalln(err)
    }
}
