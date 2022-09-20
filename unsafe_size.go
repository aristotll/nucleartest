package main

import (
    "unsafe"
    "fmt"
)

type I interface {
    XXX()
}

type S struct {
    X int64
    Y string
    C int64
}

func (s *S) XXX() {
    fmt.Println("xxx")
}

func size(i I) {
    fmt.Println(unsafe.Sizeof(i))
}

func main() {
    size(&S{1, "1", 1})
    var i uint32 = 1
    fmt.Println(unsafe.Sizeof(i))
}
