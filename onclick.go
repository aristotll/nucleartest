package main

import (
    "fmt"
)

func onclick(name string, callback func()) {
    if name == "123" {
        callback()
    }
}

func main() {
    onclick("123", func() {
        fmt.Println("onclick")
    })
}
