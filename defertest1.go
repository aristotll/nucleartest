package main

import (
    "fmt"
)

func returnButDefer() (t int) {
    defer func() {
        t = t * 10
    }()

    return 1
}

func main() {
    fmt.Println(returnButDefer())
}
