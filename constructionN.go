package main

import (
    "fmt"
    "flag"
)

var s = flag.String("s", "", "input a char")
var n = flag.Int("n", 0, "construction n char")

func main() {
    flag.Parse()
    var r string
    for i := 0; i < *n; i++ {
        r += *s
    }
    fmt.Println(r)
}
