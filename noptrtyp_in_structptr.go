package main

import (
    "fmt"
)

type csp struct {
    x int64
    y int64
}

type call struct {
    obj csp
}

func chobj(c *call) {
    if c == nil {
        return
    }
    c.obj.x = 10000
    c.obj.y = 10000
}

func main() {
    c := new(call)
    c.obj.x = 10
    c.obj.y = 10

    chobj(c)

    fmt.Println(c.obj)
}
