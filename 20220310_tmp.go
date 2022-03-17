package main

import "fmt"

func main() {
	slice := make([]int, 0, 0)
	fn(slice)
	fmt.Println("[main] slice: ", slice)
}

func fn(slice []int) {
	s := append(slice, 1)
	fmt.Println("[func] slice: ", s)
}
