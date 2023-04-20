package main

import (
    "fmt"
    "net/url"
)

func main() {
    p, err := url.JoinPath("http://www.test.com", "?")
    if err != nil {
        panic(err)
    }
    fmt.Println(p)
}
