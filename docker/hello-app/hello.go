package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!\n"))
		w.Write([]byte("Version: 1.0.0\n"))
		name, _ := os.Hostname()
		w.Write([]byte("Hostname: " + name + "\n"))
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}
