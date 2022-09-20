package main

import (
    "fmt"
)

type I interface {
    XXX()
}

type S struct {
    X int64
}

func (s *S) XXX() {}

var MAP = make(map[I]bool)

func main() {
    var s = &S{X: 123}
    var ss = &S{X: 123}
    
    MAP[s] = true
    MAP[ss] = true

    fmt.Println(MAP)
}
