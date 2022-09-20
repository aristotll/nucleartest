package main

import (
    "fmt"
)

func main() {
    fmt.Println(max("i", "love"))
}

func max(x, y string) string {
    if x > y {
        return x
    }
    return y
}
