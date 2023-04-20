package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println(os.Getenv("a"), os.Getenv("b"), os.Getenv("c"))
}
