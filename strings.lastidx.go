package main

import (
    "fmt"
    "strings"
)

func main() {
    s := "https://123/456/abc.jpeg"
    idx := strings.LastIndex(s, "/")
    fmt.Println(s[idx+1:])
}
