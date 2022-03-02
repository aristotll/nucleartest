package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	index := 3
	copy(s[index+1:], s[index:])
	fmt.Println(s[index+1:], s[index:])
	fmt.Println(s)
}
