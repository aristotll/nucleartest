package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(2)

    go func() {
        defer wg.Done()
        defer func() {
            if err := recover(); err != nil {
                fmt.Println(err)
            }
        }()        
    }()

    go func() {
        defer wg.Done()
        panic("err")
    }()

    wg.Wait()
}
