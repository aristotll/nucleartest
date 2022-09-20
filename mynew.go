package main

import (
    "fmt"
)

type S struct {}

func new() {
    fmt.Println(1)
}

func main() {
    s := new(S)
    _ = s
}
