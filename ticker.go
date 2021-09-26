package main

import (
    "fmt"
    "time"
)

func main() {
    go func() {
        ti := time.NewTicker(time.Second * 2)
        for {
            select {
            case <-ti.C:
                fmt.Println("ok")        
            }
        }
    }()
    time.Sleep(time.Second * 10)
}
