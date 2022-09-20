package main

import "fmt"

func ttt() {
	ints := make([]int, 2)
	ints = append(ints, 123)
	fmt.Println(ints)
}

func tttt() {
	var swap = func(arr []int, v1, v2 int) {
		arr[v1], arr[v2] = arr[v2], arr[v1]
	}
	arr := []int{1, 34, 66}
	fmt.Println("交换前", arr)
	swap(arr, 1, 2)
	fmt.Println("交换后", arr)
}

func main() {
	ttt()
	tttt()
}
