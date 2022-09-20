package main

import (
    "fmt"
    "flag"
)

var (
    s1 = flag.String("s1", "", "input first string")
    s2 = flag.String("s2", "", "input second string")
)

func main() {
    flag.Parse()
    max_ := max(*s1, *s2)
    min_ := ""
    if *s1 == max_ {
        min_ = *s2
    } else {
        min_ = *s1
    }
    fmt.Printf("%v < %v\n", min_, max_)
}

func max(x, y string) string {
    if x > y {
        return x
    }
    return y
}


