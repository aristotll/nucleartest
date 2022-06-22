package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i int32 = 1
	s := unsafe.Sizeof(i)
	fmt.Println(s)
}
