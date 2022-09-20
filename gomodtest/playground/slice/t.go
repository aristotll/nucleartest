package main

import (
	"fmt"
)

func Te(s *[]int, v ...int) {
	fmt.Println(v)
}

func main() {
	a := [...]int{1, 2, 3, 4, 5}
	// a1 变成了切片类型
	a1 := a[:]
	Te(&a1, a1...)

	s := make([]int, 0)
	s = append(s, 13, 3123, 312 , 141)
	fmt.Println(s)
	s = s[:len(s)-1]
	fmt.Println(s)
}
