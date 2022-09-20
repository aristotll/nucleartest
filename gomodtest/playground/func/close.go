package main

import "fmt"

// 如何把 f2 作为参数传入到 f1 中？
func f1(f func()) {
	fmt.Println(" this is f1")
	f()
}

func f2(a, b int) int {
	fmt.Println("this is f2")
	return a + b
}

func f3(f func(int, int) int, a, b int) func() {
	return func() {
		i := f(a, b)
		fmt.Println("the result is :", i)
	}
}

func main() {
	f := f3(f2, 2, 3)
	f1(f)
}

