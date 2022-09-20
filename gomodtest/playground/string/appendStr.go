package main

import (
	"fmt"
	"strings"
)

// 拼接
func appendS() {
	var strBuilder strings.Builder
	strBuilder.WriteString("")
	strBuilder.WriteString("123")
	strBuilder.WriteString("asdas")
	fmt.Println(strBuilder.String())
}

// 移除字符串第一个字母
func rmStrFirst() {
	l := make([]int, 0)
	l = append(l, 123, 55435, 534)
	fmt.Println(l)
	l = append(l[:0], l[1:]...)
	fmt.Println(l)
}

func replaceStr(s string) {
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			break
		}
		if s[i+1] <= s[i] {
			s = strings.Replace(s, string(s[i]), "", 1)
		}
	}
	fmt.Println(s)
}

func main() {
	replaceStr("2353562")
}
