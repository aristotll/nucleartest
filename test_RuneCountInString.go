package main

import (
    "unicode/utf8"
    "fmt"
)

func main() {
    s := "abc"
    s1 := "奥术大师"

    fmt.Println(utf8.RuneCountInString(s))
    fmt.Println(utf8.RuneCountInString(s1))
}
