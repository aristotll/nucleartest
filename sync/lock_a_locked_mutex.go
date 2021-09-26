package main

import (
    "sync"
    "fmt"
)

// 给已经上锁的锁继续加锁会怎样？

func main() {
    var mu sync.Mutex
    
    mu.Lock()
    defer mu.Unlock()
    fmt.Println("lock")
    
    mu.Lock()
    fmt.Println("second lock")
}
