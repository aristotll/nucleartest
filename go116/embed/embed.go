package main

import (
	"embed"
	_ "embed"
	"fmt"
	"net/http"
)

//go:embed x.txt
var txt string
//go:embed 1.jpeg
var pho embed.FS

func main() {
	http.Handle("/",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(txt)
	}))
	http.ListenAndServe(":8080", nil)
}
