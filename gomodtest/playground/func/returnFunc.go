package main

import (
	"fmt"
)

type F func(a int) int

func ret(i int) F {
	return func(a int) int {
		a = i
		return a
	}
}

func t1(f F) {
	i := f(0)
	fmt.Println(i)
}

func main() {
	f := ret(10)
	t1(f)
}
