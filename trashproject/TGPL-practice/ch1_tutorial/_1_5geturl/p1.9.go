package main

import (
	"fmt"
	"log"
	"net/http"
)

// 输出 http 的状态码
func getHttpStatus(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	status := resp.Status
	fmt.Println("the url response status is:", status)
}

func main() {
	getHttpStatus("https://www.douyu.com/9999")
}
