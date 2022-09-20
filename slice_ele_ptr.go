package main

import (
	"fmt"
)

func main() {
	a, b, c := 1, 2, 3
	s := []*int{&a, &b, &c}
	n1 := s[0]
	*n1 = 111
	for _, v := range s {
		fmt.Println(*v)
	}

	ss := []int{a, b, c}
	n2 := ss[0]
	n2 = 1
	_ = n2
	fmt.Println(ss)
}
