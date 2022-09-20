package main

import (
    "os"
    //"fmt"
)

func main() {
    _, err := os.Open("123")
    if err != nil {
        panic(err)
    }
}
