package main

import (
    "fmt"
)

type S struct {
    m map[string]string
}

func main() {
    s := new(S)
    fmt.Println(s.m["1"])
    s.m["1"] = "1"
}
