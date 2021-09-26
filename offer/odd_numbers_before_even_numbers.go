package main

import (
    "fmt"
)

// 效率低
func change_array1(arr []int) []int {
    ou := make([]int, 0)
    ji := make([]int, 0)

    for _, v := range arr {
        if v%2 == 0 {
            ou = append(ou, v)
        } else {
            ji = append(ji, v)
        }
    }

    ji = append(ji, ou[:]...)
    return ji
}

// error
func change_array2(arr []int) []int {
    p1, p2 := 0, len(arr) - 1
    for p1 < p2 {
        if arr[p1]%2 == 0 && arr[p2]%2 != 0 {
            arr[p1], arr[p2] = arr[p2], arr[p1]
            p1++
            p2--
            //fmt.Println("in swap")
        } else {
            p1++
        }
    }
    //fmt.Println(arr)
    return arr
}

func change3(nums []int) []int {
    p1, p2 := 0, len(nums)-1
    for p1 < p2 {
        if nums[p1]%2 != 0 {
            p1++
        }
        if nums[p2]%2 == 0 {
            p2--
        }
        nums[p1], nums[p2] = nums[p2], nums[p1]
    }
    return nums
}

func main() {
    a := []int{2, 4, 1, 8, 4, 3, 5}
    r := change3(a)
    fmt.Println(r)
}
