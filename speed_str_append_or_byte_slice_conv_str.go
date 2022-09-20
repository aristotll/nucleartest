package main

import (
	"fmt"
	"time"
)

func ap(s string) {
	now := time.Now()
	for i := 0; i < 255; i++ {
		s += string(byte(i))
	}

	fmt.Println("new str: ", s)
	t := time.Since(now)
	fmt.Println("字符串拼接消耗时间：", t)
}

func bs(s string) {
	now := time.Now()
	b := make([]byte, len(s))
	b = []byte(s)

	for i := 0; i < 255; i++ {
		b = append(b, byte(i))
	}

	fmt.Println("new str: ", string(b))
	t := time.Since(now)
	fmt.Println("[]byte 转 string 消耗时间：", t)
}

func main() {
	s := "123"
	ap(s)
	bs(s)
}
