package main

import "fmt"

// 可变参数的 Demo

func test(v ...int) {
	// 将多个参数一个一个的传递给 test1()
	test1(v...)
}

func test1(v ...int) {
	for _, v := range v {
		fmt.Println(v)
	}
}

func main() {
	test(1, 2, 3)
}
