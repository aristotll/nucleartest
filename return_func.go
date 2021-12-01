package main

import (
    "fmt"
)

type Func func()
type Struct struct {
    X int64
}


func fn() (*Struct, Func) {
    s := &Struct{X:1}
    return s, func() {fmt.Println(s.X)}
}

func main() {
    _, f := fn()
    f()
}
