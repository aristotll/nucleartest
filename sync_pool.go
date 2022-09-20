package main

import (
    "sync"
    "fmt"
)

type S struct {
    X int64
}

var pool = &sync.Pool{New: func() any {return &S{X: 123}} }

func a() *S {
    v := pool.Get().(*S)
    v.X = 456
    pool.Put(v)
    return v
}

func b(s *S) {
    fmt.Println(s)
}

func main() {
    b(a())
}

