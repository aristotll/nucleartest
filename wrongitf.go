package main

import (
    "fmt"
)

type S struct {A string}
type A struct {B int64}

func fn(p any) {
    fmt.Println(p.(*S).A)
}

func main() {
    fn(&A{B: 123})
}
