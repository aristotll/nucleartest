package main

import (
    "fmt"
)

func fn(s string) (string, error) {
    return s, nil
}

func main() {
    fmt.Println(fn("123"))
}
