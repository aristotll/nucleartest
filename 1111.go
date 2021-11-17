package main

import (
    "fmt"
)

func t1111() int {
    n := []int{1, 1, 2, 3, 3, 5, 5}
    r := 0

    for i := 0; i < len(n); i++ {
        r ^= n[i]
    }

    return r
}

func main() {
    fmt.Println(t1111())
}
