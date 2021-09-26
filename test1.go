package main

import (
    "fmt"
    "time"
    //"sync"
    "sync/atomic"
)

func main() {
    var x int64
    //var mux sync.Mutex
    go func() {
        for {
            atomic.AddInt64(&x, 1)
            //mux.Lock()
            x++
            //mux.Unlock()
        }
    }()
    time.Sleep(time.Second)
    fmt.Println("x = ", x)
}
