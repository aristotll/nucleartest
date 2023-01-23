package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

var text = `
asdsaijilfgilfdfdg
werwer
kljgklhjkllewkrioweut
xcm,n,m234234sdf
`

func main() {
	addr := "localhost.proxyman.io:8081"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fluster := w.(http.Flusher)
		reader := bufio.NewReader(strings.NewReader(text))
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			}
			fmt.Fprint(w, line)
			fluster.Flush()
			<-time.Tick(time.Second)
		}
	})
	log.Printf("server listen in %v\n", addr)
	log.Fatalln(http.ListenAndServe(addr, nil))
}
