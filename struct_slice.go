package main

import (
    "fmt"
)

type st struct {
    b []int
}

func f(s st) {
    s.b[0] = 123
}

func main() {
    s := st{
        b: []int{1, 2, 3},
    }
    f(s)
    fmt.Println(s)
}
