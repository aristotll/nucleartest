package main

import (
    "fmt"
    "unsafe"
)

type S struct {
    A int64
    B int32
    C string
}

func main() {
	i := 100
	fmt.Printf("%p \n", &i)
	fmt.Println(uintptr(i))

    var new uintptr = 2
    fmt.Println(new)
    fmt.Println((*S)(unsafe.Pointer(new)))
}
