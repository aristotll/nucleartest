package main

import (
	"fmt"
)

type A struct {
	s string
}

func main() {
	var a *A
	if check(a) {
		a, err := generate()
		fmt.Println(a.s, err)
	}
	fmt.Println(a.s)
}
func generate() (*A, error) {
	return &A{s: "b"}, nil
}
func check(a *A) bool {
	return true
}
