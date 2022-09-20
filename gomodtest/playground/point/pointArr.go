package main

import "fmt"

func t(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		fmt.Println((*arr)[i])
	}
}

func main() {
	arr := make([]int, 10)
	arr = []int{213, 344435, 346, 54, 6, 4}
	t(&arr)
}