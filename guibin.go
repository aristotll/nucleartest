package main

import (
    "fmt"
)

func main() {
    v := sort([]int{2, 4, -1, 9, 0})
    fmt.Println(v)
}

func sort(n []int) []int {
    if len(n) < 2 {
        return n
    }
    m := len(n) >> 1
    left := sort(n[:m])
    right := sort(n[m:])
    return merge(left, right)
}

func merge(x, y []int) []int {
    tmp := make([]int, 0)
    i, j := 0, 0
    for i < len(x) && j < len(y) {
        if x[i] < y[j] {
            tmp = append(tmp, x[i])
            i++
        } else {
            tmp = append(tmp, y[j])
            j++
        }
    }
    if i < len(x) {
        tmp = append(tmp, x[i:]...)
    }
    if j < len(y) {
        tmp = append(tmp, y[j:]...)
    }
    return tmp
}
