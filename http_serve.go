package main

import (
	"net"
	"net/http"
    "os"
    "log"
)

func main() {
    addr := os.Getenv("addr")
    if addr == "" {
        addr = ":6666"
    }
    log.Printf("listen in %v\n", addr)
	l, e := net.Listen("tcp", addr)
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
