package main

import "fmt"

func test(arr *[]int)  {
	a := *arr
	a[0] = 1111
	fmt.Printf("func 中 arr 的地址 %p", arr)
}

func test1(a *int) {
	i := &a
	*a = 100
	fmt.Printf("func %p\n", i)
	fmt.Println(*a)
}

func main() {
	//arr := make([]int, 10)
	//test(&arr)
	//fmt.Println(arr[0])
	//
	//fmt.Printf("main 中 arr 的地址 %p\n", arr)

	a := 123
	test1(&a)
	fmt.Printf("main %p:\n", &a)
	fmt.Println(a)
}
