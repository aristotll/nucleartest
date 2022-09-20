package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

/*
	找一个产生大量数据的网站。连续两次运行 fetch，看报告的时间是否会有大的变化，调查缓存情况。
	每一次获取的内容一样吗？修改 fetch 将内容输出到文件，这样可以检查他是否一致
*/

func fetch1(url string, ch chan string) {
	start := time.Now()

	// get url
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("get url error: %s", err)
		return
	}

	// 准备一个 txt 用来写入 response
	file, err := os.OpenFile("ch1_tutorial/_1_6geturlconcurrent/response.txt",
		os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		ch <- fmt.Sprintf("operating file error: %s", err)
		return
	}

	// write response body to file
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("read response body error: %s", err)
		return
	}
	// 美化一下 response，多次写入后辨别更直观
	reader := strings.NewReader(string(bytes) + "\n\n\n"+
		"============================ 【end there】 ==============================\n\n")

	n, err := io.Copy(file, reader)
	if err != nil {
		ch <- fmt.Sprintf("copy error: %s", err)
		return
	}

	defer resp.Body.Close()
	defer file.Close()

	end := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %d %s", end, n, url)
}

func main() {
	ch := make(chan string)
	for i := 0; i < 5; i++ {
		go fetch1("https://www.douyu.com/9999", ch)
	}
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}

	// Output first
	// 0.30s 101638 https://www.douyu.com/9999
	// 0.60s 101638 https://www.douyu.com/9999
	// 0.75s 101638 https://www.douyu.com/9999
	// 0.75s 101638 https://www.douyu.com/9999
	// 0.75s 101638 https://www.douyu.com/9999

	// Output second
	// 0.59s 101696 https://www.douyu.com/9999
	// 0.59s 101696 https://www.douyu.com/9999
	// 0.60s 101696 https://www.douyu.com/9999
	// 0.60s 101696 https://www.douyu.com/9999
	// 0.61s 101696 https://www.douyu.com/9999
}
