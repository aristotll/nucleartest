package main

import (
	"fmt"
)

type S struct {
	A int64
	B string
}

func main() {
	s := S{10, "abc"}
	c := make(chan S, 1)
	c <- s
	ss := <-c
	ss.A = 999
	fmt.Println("ss: ", ss)
	fmt.Println("s: ", s)
}
