package main

import (
	"net/http"
	"log"
)

func main() {
	if err := http.ListenAndServe("http://localhost:9999", nil); err != nil {
		log.Fatalln(err)
	}
}
