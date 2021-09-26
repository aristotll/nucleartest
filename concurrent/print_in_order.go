package main

import (
    "fmt"
    //"sync"
)

//var wg sync.WaitGroup
var done, c2, c3 chan struct{}

type gors struct {
    A, B, C func()
}

func new(f1 func(), f2 func(), f3 func()) *gors {
    return &gors{
        A: f1,
        B: f2,
        C: f3,
    }
}

func (g *gors) Start() {
    go func() {
        //defer wg.Done()
        g.A()
    }()

    go func() {
        //defer wg.Done()
        g.B()
    }()

    go func() {
        //defer wg.Done()
        g.C()
    }()
}

func main() {
    f1 := func() {
        fmt.Println("First")
        c2 <- struct{}{}
    }
    f2 := func() {
        <-c2
        fmt.Println("Second")
        c3 <- struct{}{}
    }
    f3 := func() {
        <-c3
        fmt.Println("Thrid")
        done <- struct{}{}
    }
    gs := new(f1, f2, f3)
    //wg.Add(3)
    gs.Start()
    //wg.Wait()
    //<-c3
    <-done
}
