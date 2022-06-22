package main

import (
	"fmt"
)

func appendSliPtr(s *[]int) {
	*s = append(*s, 1, 2, 3)
}

func appendSli(s []int) {
	s = append(s, 999)
}

func main() {
	s := make([]int, 0, 30)
	appendSliPtr(&s)
	fmt.Println(s)

	s = []int{1, 1, 1}
	appendSli(s)
	fmt.Println(s)
}
