package main

import (
    "fmt"
)

func main() {
    var m = make(map[int]map[int]int)
    _, ok := m[1]
    if !ok {
        fmt.Println("!ok")
        m[1] = make(map[int]int)
    }
    m[1][1] = 1
    fmt.Println(m)
}
