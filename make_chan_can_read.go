package main

import (
    "fmt"
    "time"
)

func main() {
    c := make(chan struct{})
    
    select {
    case <-c:
        fmt.Println("ok")
    }

    //go func() {
    //    <-c
    //    fmt.Println("ok")
    //}()

    time.Sleep(time.Second * 3)
}
