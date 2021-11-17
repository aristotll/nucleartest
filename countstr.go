package main

import (
    "fmt"
    "flag"
    "unicode/utf8"
)

var s = flag.String("s", "", "input a string")

func main() {
    flag.Parse()
    fmt.Println(utf8.RuneCountInString(*s))
}
