package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := "123"
	b := []byte(s)

	fmt.Printf("%p %p \n", &s, &b)

	bb := *(*[]byte)(unsafe.Pointer(&s))
	bb = append(bb, '8')
	fmt.Printf("%p %p \n", &s, &bb)
	fmt.Println(s, b)
}
