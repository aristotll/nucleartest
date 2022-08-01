package main

import "fmt"

type fn func(int, int) int

type S1 struct {}

type Struct struct {
	fn
	S1
}

func (f fn) Do(x, y int) int {
	return f(x, y)
}

func (s *S1) Do1() {
	fmt.Println("s1 func")
}

func main() {
	s := &Struct{
		fn: func(x, y int) int {
			return x + y
		},
	}
	fmt.Println(s.Do(1, 2))
	s.Do1()
}
