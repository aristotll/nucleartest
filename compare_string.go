package main

import (
    "fmt"
    "flag"
)

var x = flag.String("x", "", "input string1")
var y = flag.String("y", "", "input string2")

func main() {
    flag.Parse()
    fmt.Printf("%v is larger in %v and %v\n", max(*x, *y), *x, *y)
}

func max(x, y string) string {
    if x > y {
        return x
    }
    return y
}
