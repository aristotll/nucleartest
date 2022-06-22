package main

import (
	"fmt"
)

func f(x *int64) {
	fmt.Printf("x: %p, &x: %p, *x: %v\n", x, &x, *x)
}

func main() {
	var a int64 = 5
	fmt.Printf("&a: %p\n", &a)
	b := &a
	fmt.Printf("b: %p, &b: %p\n", b, &b)
	f(b)
}
