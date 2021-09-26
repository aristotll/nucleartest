package main

import (
    "fmt"
)

func ch(n []int) {
    //nn := make([]int, 0)
    //for _, v := range n {
    //    nn = append(nn, v+1)
    //}
    //for i := 0; i < len(n); i++ {
    //    nn = append(nn, n[i]+1)
    //}
    //fmt.Println("func: ", nn)
    //n = nn
    n[0] = 888
}

func chp(n *[]int) {
    nn := make([]int, 0)
    for _, v := range *n {
        nn = append(nn, v+1)
    }
    *n = nn
}

func main() {
    n := []int{1, 2, 3, 4, 5}
    chp(&n)
    fmt.Println(n)
}
