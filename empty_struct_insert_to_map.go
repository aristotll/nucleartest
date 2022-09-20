package main

import (
    "fmt"
)

type S struct {}

type SS struct {
    X int64
}

func main() {
    s1, s2 := &S{}, &S{}
    m := make(map[*S]struct{})
    m[s1] = struct{}{}
    m[s2] = struct{}{}
    fmt.Println(m)

    s11, s22 := &SS{1}, &SS{2}
    mm := make(map[*SS]struct{})
    mm[s11] = struct{}{}
    mm[s22] = struct{}{}
    fmt.Println(mm)
}
