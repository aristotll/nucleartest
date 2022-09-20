package main

import (
    "fmt"
)

func main() {
    s1 := "abc"
    s2 := "abcd"
    compare(s1, s2)

    s1 = "abc"
    s2 = "def"
    compare(s1, s2)

    s1 = "abc"
    s2 = "aBc"
    compare(s1, s2)
}

func compare(s1, s2 string) {
    if s1 > s2 {
        fmt.Printf("%s > %s\n", s1, s2)
    } else {
        fmt.Printf("%s > %s\n", s2, s1)
    }
}
