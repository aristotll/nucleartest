package main

import (
	"fmt"
)

func main() {
	x := 5
	y := x
	fmt.Printf("x: %d, y: %d \n", x, y)

	y = 10
	fmt.Printf("x: %d, y: %d \n", x, y)

	s := "123"
	s1 := s
	fmt.Printf("s: %s, s1: %s \n", s, s1)

	s1 = "456"
	fmt.Printf("s: %s, s1: %s \n", s, s1)

}
