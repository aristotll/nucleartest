package main

import (
	"fmt"
)

func f(s []byte, i, j int) {
	if i == j {
		return
	}
	s[i], s[j] = s[j], s[i]
	f(s, i+1, j-1)
}

func main() {
	var s = []byte("abcdefg")
	f(s, 0, len(s)-1)
	fmt.Println(string(s))
}
