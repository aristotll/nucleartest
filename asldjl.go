package main

import (
    "fmt"
)

type S struct {A string}

func main() {
    vec := []*S{
        &S{A: "123"},
    }
    fmt.Println(vec)
}
