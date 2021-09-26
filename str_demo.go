package main

import (
    "fmt"
)

func main() {
    s1 := "123"
    s2 := "123"
    fmt.Println(s1 == s2)

    sp1 := &s1
    sp2 := &s2
    fmt.Println(sp1 == sp2)
}
