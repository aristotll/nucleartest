package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Set("Content-Type", "json")
		//w.Header().Set("X-Content-Type-Options", "nosniff")
		http.ServeFile(w, r, "./demo.lubenwei")
	})

	log.Println("listen in 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
