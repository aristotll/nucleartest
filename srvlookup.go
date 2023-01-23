package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
)

const dataFile = "/var/data/kubia.txt"
const serviceName = "kubia.default.svc.cluster.local"
const port = 8080

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			b, err := io.ReadAll(r.Body)
			defer r.Body.Close()
			if err != nil {
				writeErrMsg(w, err)
				return
			}
			if err := os.WriteFile(dataFile, b, 0644); err != nil {
				writeErrMsg(w, err)
				return
			}
			w.WriteHeader(http.StatusOK)
		case "GET":
			w.WriteHeader(http.StatusOK)
			if r.URL.String() == "/data" {
				b, err := os.ReadFile(dataFile)
				if err != nil {
					writeErrMsg(w, err)
					return
				}
				w.Write(b)
			} else {
				hostname, _ := os.Hostname()
				w.Write([]byte(fmt.Sprintf("You've hit %v \n", hostname)))
				w.Write([]byte("Data stored in the cluster:\n"))
				_, addrs, err := net.LookupSRV("http", "tcp", serviceName)
				if err != nil {
					writeErrMsg(w, err)
					return
				}
				for _, addr := range addrs {
					url := fmt.Sprintf("http://%v:%v/data", addr.Target, addr.Port)
					req, err := http.NewRequest("GET", url, nil)
					if err != nil {
						writeErrMsg(w, err)
						return
					}
					resp, err := http.DefaultClient.Do(req)
					if err != nil {
						writeErrMsg(w, err)
						return
					}
					b, err := io.ReadAll(resp.Body)
					if err != nil {
						writeErrMsg(w, err)
						return
					}
					fmt.Appendf([]byte("-", addr.))
					w.Write(b)
				}
			}
		}
	})
	addr := fmt.Sprintf(":%v", port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}

func writeErrMsg(w http.ResponseWriter, err error) {
	log.Println(err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}
