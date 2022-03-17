package main

import (
	"reflect"
	"unsafe"
	"fmt"
)

func getSliceHeader(slice []int) {
	//n := []int{1, 2, 3}
	struc := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	fmt.Printf("%+v\n", struc)

	dataPtr := unsafe.Pointer(struc.Data)
	fmt.Println(dataPtr)
	
	//data := *(*[3]int)(dataPtr)
	//fmt.Println(data)
}

func sliceAppend(slice []int) {
	slice = append(slice, 1, 2, 3, 4)
	fmt.Println("[func] sliceHeader info: ")
	getSliceHeader(slice)
	fmt.Println("========================")

	struc := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	dataPtr := unsafe.Pointer(struc.Data)
	data := *(*[4]int)(dataPtr)
	fmt.Println("[func] after append, sliceHeader.data: ", data)
}

func main() {
	n := make([]int, 0, 20)
	fmt.Println("[main] sliceHeader info: ")
	getSliceHeader(n)
	fmt.Println("========================")

	sliceAppend(n)
	fmt.Println("[main] after call, slice data: ", n)
	struc := (*reflect.SliceHeader)(unsafe.Pointer(&n))
	dataPtr := unsafe.Pointer(struc.Data)
	data := *(*[4]int)(dataPtr)
	fmt.Println("[main] sliceHeader.data", data)
}