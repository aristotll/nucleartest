package main

import (
	"fmt"
	"net/http"
	"time"
)

func api(callname string, w http.ResponseWriter, callbackFunc func(w http.ResponseWriter)) {
	fmt.Printf("run %v ...\n", callname)
	w.Write([]byte("执行中..."))
	time.Sleep(time.Second * 5)
	// 执行完毕，调用回调函数通知调用方
	callbackFunc(w)

}

func main() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		api("testgeneric", w, func(w http.ResponseWriter) {
			w.Write([]byte("执行完成"))
		})
	})
	http.ListenAndServe(":9527", nil)
}
