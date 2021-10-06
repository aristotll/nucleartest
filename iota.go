package main

import (
    "fmt"
)

const (
    one = 1 << iota
    two
	three = iota
)

const (
	a = iota
	b
	c
	d = iota
	e = iota
	aa = 100
	bb = iota
	cc
)

func main() {
    fmt.Println(one, two, three)
	fmt.Println(a, b, c, d, e, aa, bb, cc)
}
