package main

import (
    "fmt"
)

type TT struct {
    name string
}

func tslice() []TT {
    n := make([]*TT, 1, 1)
    t := &TT{}
    n[0] = t

    nn := make([]TT, 1, 1)
    t1 := TT{}
    nn[0] = t1

    return nn
}

func tnew() {
    t := new(TT)
    _ = t
}

func main() {
    n := tslice()
    fmt.Println(n)
    tnew()
}
