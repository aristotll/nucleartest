package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    ch := make(chan int, 100)
    var wg sync.WaitGroup
    wg.Add(1)

    go func() {
        defer wg.Done()
        for {
            time.Sleep(time.Second*10)
            ch <- 1
        }
    }()

    wg.Add(1)
    go func() {
        defer wg.Done()
        for {
            time.Sleep(time.Second*10)
            fmt.Println(<-ch)        
        }
    }()

    wg.Wait()
}
