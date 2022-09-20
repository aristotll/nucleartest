package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

/*
	为传入的 url 提供一个 http:// 前缀
*/

func AddPrefix(url string) {
	var sb strings.Builder
	prefix := "http://"
	if !strings.HasPrefix(url, prefix) {
		sb.WriteString(prefix)
		sb.WriteString(url)
	}else {
		sb.WriteString(url)
	}

	resp, err := http.Get(sb.String())
	if err != nil {
		log.Fatal(err)
	}

	readAll, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("url: ", sb.String())
	fmt.Printf("%s", readAll)
}

func main() {
	AddPrefix("www.douyu.com/9999")
}
