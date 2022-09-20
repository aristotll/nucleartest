package main

import (
	"fmt"
)

func main() {
	s := make([]int, 0, 10)
	s = s[0 : cap(s)-1 : 1]
	fmt.Println(s)
}
