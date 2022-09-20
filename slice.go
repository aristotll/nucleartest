package main

import (
	"fmt"
)

func sliceT(a []int) {
	ri := 2
	// 有效
	a[0] = 111
	fmt.Printf("a address: %p\n", &a)
	a = append(a[:ri], a[ri+1:]...)
	fmt.Printf("after remove, a address: %p\n", a)
	// 下面的全部无效
	//a = append(a, 1)
	//fmt.Printf("after append, a address: %p\n", a)
	//a[0] = 1
	//a = append(a[:ri], a[ri+1:]...)
}

func slicePoint(a *[]int) {
	*a = append(*a, 1)
	(*a)[0] = 111
	ri := 2
	*a = append((*a)[:ri], (*a)[ri+1:]...)
}

func main() {
	a := []int{1, 2, 3, 4}
	fmt.Println("before: ", a)
	sliceT(a)
	fmt.Println("after :", a)
	slicePoint(&a)
	fmt.Println("after point: ", a)
}
