package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// const dataFile = "/var/data/kubia.txt"
var dataFile = filepath.Join(getHome(), "/pj/justtest/kubia.txt")

func getHome() string {
	return os.Getenv("HOME")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			f, err := os.OpenFile(dataFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
			if checkErrorAndWriteResponse(w, "open file error: %v\n", err) {
				return
			}
			defer f.Close()

			_, err = io.Copy(f, r.Body)
			if checkErrorAndWriteResponse(w, "io.Copy error: %v\n", err) {
				return
			}
		} else {
			f, err := os.Open(dataFile)
			if checkErrorAndWriteResponse(w, "open file error: %v\n", err) {
				return
			}
			defer f.Close()

			w.WriteHeader(http.StatusOK)
			hostname, _ := os.Hostname()
			io.WriteString(w, "You've hit "+hostname+"\n")
			io.WriteString(w, "Data stored on this pod: ")
			_, err = io.Copy(w, f)
			if checkErrorAndWriteResponse(w, "io.Copy error: %v\n", err) {
				return
			}
			io.WriteString(w, "\n")
		}
	})
	if err := http.ListenAndServe(":7788", nil); err != nil {
		log.Fatalln(err)
	}
}

func checkErrorAndWriteResponse(w http.ResponseWriter, formart string, err error) (hasErr bool) {
	if err == nil {
		hasErr = false
		return
	}
	hasErr = true
	errstr := fmt.Sprintf(formart, err)
	log.Println(errstr)
	w.WriteHeader(http.StatusInternalServerError)
	io.WriteString(w, errstr)
	return
}
