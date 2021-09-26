package main

import (
	"fmt"
)

// 从第一个字符串中，删除所有出现在第二个字符串中的字符
// 例如 s1 => "hello" s2 => "el" result => "ho"

func remove(s1, s2 string) {
	//m := make(map[rune]bool)
	n := [26]int{}
	var newStr string
	for _, char := range s2 {
		n[char-'a']++
	}

	for _, char := range s1 {
		if n[char-'a'] == 0 {
			newStr += string(char)
		}
	}

	fmt.Println(newStr)
}

func main() {
	remove("hello", "el")
}
