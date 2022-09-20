package main

import (
    "fmt"
)

type St struct {
    m *map[string]string
}

func (s *St) Push(k, v string) {
    (*s.m)[k] = v
}

func main() {
    s := &St{}
    s.m = &map[string]string{}
    s.Push("name", "zhang")
    fmt.Println(s.m)
}
