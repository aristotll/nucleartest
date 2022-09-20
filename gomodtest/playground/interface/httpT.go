package main

import (
	"fmt"
	"net/http"
)

type dollars float64

type database map[string]dollars

func (d database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for item, price := range d {
		_, _ = fmt.Fprintf(w, "%s: %f \n", item, price)
	}
}

func main() {
	db := database{"shoes": 500, "socks": 5}
	http.ListenAndServe(":8080", db)
}
