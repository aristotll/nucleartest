package main

import (
	"log"
	"net/http"
)

func main() {
	if err := http.ListenAndServe("http://localhost:9999", nil); err != nil {
		log.Fatalln(err)
	}
}
