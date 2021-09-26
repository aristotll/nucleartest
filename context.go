package main

import (
    "fmt"
    "context"
    "time"
)

func main() {
    _, cancel := context.WithTimeout(context.Background(), time.Second*3)
    defer cancel()

    for {
        fmt.Println("123")
        time.Sleep(time.Second)
    }
}
