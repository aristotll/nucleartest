package main

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func getUrl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	r, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	r = strings.Replace(r, "\n", "", -1)
	getUrl(r)
}
