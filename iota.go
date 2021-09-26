package main

import (
    "fmt"
)

const (
    one = 1 << iota
    two
)

func main() {
    fmt.Println(one, two)
}
