package main

import (
    "fmt"
)

func f(s []int) {
    ss := []int{1, 2, 3, 4, 5}
    copy(s, ss)
}

func main() {
    s := make([]int, 10, 10)
    f(s)
    fmt.Println(s)
}
