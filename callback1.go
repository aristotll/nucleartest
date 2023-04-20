package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)
    ch3 := make(chan int)

    go time.AfterFunc(time.Second*3, func() {
        fmt.Println("after 3000[1]")
        ch1 <- 1
    })
    fmt.Println("hello[1]")
    <-ch1

    go time.AfterFunc(time.Second*3, func() {
        fmt.Println("after 3000[2]")
        ch2 <- 1
    })
    fmt.Println("hello[2]")
    <-ch2

    go time.AfterFunc(time.Second*3, func() {
        fmt.Println("after 3000[3]")
        ch3 <- 1
    })
    fmt.Println("hello[3]")
    <-ch3
}
