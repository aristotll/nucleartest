package main

import (
	"fmt"
)

func main() {
	s := "aaa"
	var v int = int(s[0])

	for i := 1; i < len(s); i++ {
		v ^= int(s[i])
	}

	fmt.Println(v)
}
