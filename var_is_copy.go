package main

import "fmt"

type AAA struct {
	A int64
}

func main() {
	i := []int{1, 2, 3}
	ii := i
	ii = append(ii, 5)
	fmt.Println(i, ii)

	s := AAA{A: 123}
	ss := s
	ss.A = 456
	fmt.Println(s, ss)

	a := 1
	b := a
	b = 123
	fmt.Println(a, b)
}
