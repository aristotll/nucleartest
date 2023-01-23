package main

import (
	"net"
	"net/http"
)

func main() {
	l, e := net.Listen("tcp", ":6666")
	if e != nil {
		panic(e)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello, world\n"))
	})

	if err := http.Serve(l, nil); err != nil {
		panic(err)
	}
}
