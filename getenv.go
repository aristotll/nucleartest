package main

import (
    "os"
    "fmt"
)

func main() {
    whoami := os.Getenv("HOME")
    fmt.Println(whoami)
}
