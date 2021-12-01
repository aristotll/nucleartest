package main

import (
    "time"
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(1)
    
    go func() {
        ticker := time.NewTicker(time.Second)
        defer wg.Done()
        for {
            <-ticker.C
            fmt.Println("tick!")
        }
    }()
    
    wg.Wait()

}
