package main

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	addr := ":8080"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.Copy(w, strings.NewReader("Hello"))
	})
	log.Printf("listen in %v \n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalln(err)
	}
}
