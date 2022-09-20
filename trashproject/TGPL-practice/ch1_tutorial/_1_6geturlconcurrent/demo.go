package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

/*
	并发获取很多 url 内容
 */

func fetch(url string, ch chan string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	n, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v ", url, err)
		return
	}

	end := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs: %d %s", end, n, url)
}

func main() {
	start := time.Now()
	urls := make([]string, 0)
	urls = append(urls, "https://www.douyu.com/9999",
			"https://www.douyu.com/610588",
			"https://search.bilibili.com/all?keyword=数据结构&from_source=nav_suggest_new",
			"https://translate.google.cn/#view=home&op=translate&sl=auto&tl=en&text=并发")
	ch := make(chan string)

	for _, url := range urls {
		go fetch(url, ch)
	}
	for range urls {
		fmt.Println(<-ch)
	}
	fmt.Printf("%f elapsed\n", time.Since(start).Seconds())
}

