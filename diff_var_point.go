package main

import (
    "fmt"
)

type Struct struct {
    X, Y int64
}

func (i *Struct) String() string {
    return fmt.Sprintf("X: %d, Y: %d", i.X, i.Y)
}

func main() {
    var s *Struct
    s1 := &Struct{}

    fmt.Println(s, s1)
    fmt.Println(s.String())
    fmt.Println(s1)
}
