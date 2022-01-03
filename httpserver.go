package main

import (
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("123"))
	})

	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatalln(err)
	}
}
