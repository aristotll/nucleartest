package main

import (
//	"bufio"
	"fmt"
	"net/http"
)

func chunkHandler(w http.ResponseWriter, r *http.Request) {
	// 设置响应头，表示使用 chunked 编码
	w.Header().Set("Transfer-Encoding", "chunked")

	// 将响应写入 chunked 格式
	fmt.Fprintf(w, "5\r\nHello\r\n")
	fmt.Fprintf(w, "5\r\nWorld\r\n")
	fmt.Fprintf(w, "0\r\n\r\n")
}

func main() {
	http.HandleFunc("/chunked", chunkHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
