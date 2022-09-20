package main

import (
    "fmt"
)

type T []int

func main() {
    t := &T{}
    *t = append(*t, 1)
    fmt.Println((*t)[0])
}
