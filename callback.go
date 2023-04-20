package main

import (
    "time"
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(3)
    go time.AfterFunc(time.Second*3, func() {
        defer wg.Done()
        fmt.Println("after 3000[1]")
        go time.AfterFunc(time.Second*3, func() {
            defer wg.Done()
            fmt.Println("after 3000[2]")
            go time.AfterFunc(time.Second*3, func() {
                defer wg.Done()
                fmt.Println("after 3000[3]")
            })
            fmt.Println("hello[3]")
        })
        fmt.Println("hello[2]")
    })
    fmt.Println("hello[1]")

    wg.Wait()
}
