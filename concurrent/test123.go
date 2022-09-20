package main

import (
    "fmt"
)

func fn(n1, n2 []int) []int {
    newn := make([]int, 0, len(n1)+len(n2))
    n1l := len(n1)
    n2l := len(n2)
    i := 0
    j := 0

    for i < n1l && j < n2l {
        if n1[i] > n2[j] {
            newn = append(newn, n2[j])
            j++
        } else {
            newn = append(newn, n1[i])
            i++
        }
    }

    if i < n1l {
        newn = append(newn, n1[i:]...)
    }
    if j < n2l {
        newn = append(newn, n2[j:]...)
    }

    return newn
}

func main() {
    s1 := []int{1, 2, 3, 5}
    s2 := []int{4, 5, 6, 8}
    fmt.Println(fn(s1, s2))
} 
