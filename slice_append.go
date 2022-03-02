package main

import "fmt"

func main() {
	slice := make([]int, 0, 10)
	slice = fn(slice)
	//fmt.Printf("main point: %p\n", &slice)
	fmt.Println(slice)

	fnPtr(&slice)
	fmt.Println(slice)
}

func fn(slice []int) (s []int) {
	//fmt.Printf("before point: %p\n", slice)
	s = append(slice, 1)
	return
	//slice = s
	//fmt.Println(cap(slice))
	//fmt.Printf("after append point: %p\n", slice)
}

func fnPtr(slicePtr *[]int) {
	*slicePtr = append(*slicePtr, 1)
}
