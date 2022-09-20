package main

import (
	"fmt"
)

func reverse(s *[]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
	fmt.Println(*s)
}

func reverse1(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	fmt.Println(s)
}

func main() {
	a := []int{123, 324, 453, 657, 87, 15}
	reverse(&a)
	reverse1(a)
}
