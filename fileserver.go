package main

import (
	"log"
	"net/http"
)

func main() {
	if err := http.ListenAndServe(":8081", http.FileServer(http.Dir("/Volumes/4t/"))); err != nil {
		log.Fatalln(err)
	}

}
