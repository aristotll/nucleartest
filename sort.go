package main

import (
    "fmt"
    "sort"
)

func main() {
    arr := []int{3, 2, 1, 5, 6, 4}
    fmt.Println("排序前：", arr)
    sort.Slice(arr, func(i, j int) bool {
        return arr[i] < arr[j]
    })
    fmt.Println("排序后：", arr)
}
