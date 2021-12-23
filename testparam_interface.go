package main

import "fmt"

func fn(x interface{}) {
	fmt.Println(x)
}

func main() {
	fn(1)
	fn(2)

	var i interface{}
	i = 123
	i = "1"

	fmt.Println(i)
}