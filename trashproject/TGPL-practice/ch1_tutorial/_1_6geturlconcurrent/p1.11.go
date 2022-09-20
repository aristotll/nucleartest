package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

/*
	使用更长的参数列表尝试 fetch，例如使用 alexa.com 排名前 100 万的网站。
	如果一个网站没有响应，程序的行为是怎样的？
	【不会写】
*/

var wg sync.WaitGroup

func fetch2(ctx context.Context, url string, ch chan string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("get url error: %s", err)
		return
	}
	
	n, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("copy error: %s", err)
		return
	}
	end := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %d %s", end, n, url)
	defer wg.Done()
}

func main() {
	// 超过 5 秒，判定为超时
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancelFunc()

	// 神秘网站，访问会超时
	url := "https://www.google.com.hk"
	_ = url
	// 能访问的网站
	trueURL := "https://www.baidu.com"
	_ = trueURL

	ch := make(chan string, 5)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go fetch2(ctx, trueURL, ch)
		fmt.Println("ch <-")
	}


	for {
		select {
		case <-ctx.Done():
			fmt.Println("get url time out")
		case <-ch:
			fmt.Println("<===========")
			fmt.Println(<-ch)
		default:
			<-time.Tick(time.Second * 3)
			fmt.Println("default")
		}
	}
	wg.Wait()

}
