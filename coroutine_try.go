package main

import (
    "fmt"
)

const count = 4

var (
    chans = func(count int) (r []chan int) {
        for i := 0; i < count; i++ {
            r = append(r, make(chan int, 1))
        }
        return r
    }(count)
)

func printValue(value int) {
    select {
    case <-chans[0]:
        fmt.Println(value)
    case <-chans[1]:
        fmt.Println(value + 1)
    case <-chans[2]:
        fmt.Println(value + 2)
    case <-chans[3]: 
        fmt.Println(value + 3)
    }
}

func main() {
    for i := 0; i < count; i++ {
        chans[i] <- 1
        printValue(1)
    }
}
