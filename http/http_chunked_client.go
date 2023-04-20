package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	// 创建请求
	req, err := http.NewRequest("GET", "http://localhost:8080/chunked", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// 读取 chunked 响应
	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		// 解析 chunk size
		chunkSize, _ := strconv.ParseInt(strings.TrimSpace(string(line)), 16, 64)
		if chunkSize == 0 {
			break
		}

		// 读取 chunk 数据
		chunk := make([]byte, chunkSize)
		_, err = reader.Read(chunk)
		if err != nil {
			fmt.Println(err)
			return
		}

		// 读取结束符
		_, err = reader.ReadBytes('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print(string(chunk))
	}
}
