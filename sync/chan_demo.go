package main

import (
    "fmt"
    "time"
)

var c = make(chan int)

func main() {
    go func() {
        c <- 1
    }()

    go func() {
        fmt.Println(<-c)
    }()

    time.Sleep(time.Second * 3)
}
