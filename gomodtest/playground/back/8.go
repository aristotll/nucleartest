package main

import (
	"fmt"
)

func main() {
	m := make(map[int]int)
	m[1] = 2

	if v, ok := m[1]; ok {
		fmt.Println(v)
	}
}
