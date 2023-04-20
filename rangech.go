package main

import (
    "fmt"
    "time"
)

func main() {
    var ch = make(chan int, 1000)
    var i int = 1

    go func() {
        time.Sleep(time.Second*5)
        for {
            ch <- i
            i++
            time.Sleep(time.Second)
        }
    }()

    fmt.Println(len(ch))

    for {
        select {
        case v := <- ch:
            fmt.Println(v)
        }
    }

    //time.Sleep(time.Second*10)
    select{}
}
