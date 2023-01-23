package main

import (
    "fmt"
)

var b = make([]byte, 1024)

func main() {
    set(7)
}

func set(n int) {
    sliceIndex := n / 8
    offsetIndex := n % 8
    b[sliceIndex] |= 1 << offsetIndex
    fmt.Println(b[sliceIndex])
}
