package main

import (
    "fmt"
)

const (
    __RDWR__ = 0x0002
    __APPEND__ = 0x0008
    __CREAT__ = 0x0200
)

func main() {
    fmt.Printf("%x\n", __RDWR__|__APPEND__|__CREAT__)
}
