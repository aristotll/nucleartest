package main

import (
    "fmt"
    "strings"
)

func main() {
    s := "123.jpg?KID=123"
    idx := strings.Index(s, "?")
    s = s[:idx+1]
    fmt.Println(s)
}
