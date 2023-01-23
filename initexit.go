package main

import (
    "os"
    "fmt"
)

func init() {
    fmt.Println(1)
    os.Exit(0)
}

func main() {
    fmt.Println("123")
}
