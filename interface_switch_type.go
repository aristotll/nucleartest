package main

import (
	"fmt"
)

type i interface {
	f(int, int) int
}

type s struct{}

func (s) f(x, y int) int {
	return x + y
}

func f() s {
	return s{}
}

func main() {
	// 搞错了，还以为是强转，其实只是调用方法
	fmt.Println(f().f(1, 2))
}
