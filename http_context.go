package main

import (
	"context"
	"log"
	"net/http"
)

func request() {
	ctx := context.WithValue(context.Background(), "user", "123")
	r, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:11111", nil)
	if err != nil {
		panic(err)
	}
	r.WithContext()
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func main() {
	var startServer = func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			u := r.Context().Value("user")
			log.Println(u)
		})
		http.ListenAndServe(":11111", nil)
	}
	go startServer()
	request()
	select {}
}
