package main

import (
	"fmt"
)

func findMin(arr []int) int {
	left := 0
	right := len(arr) - 1

	for left < right {
		mid := (left + right) / 2
        // 属于第一个子数组
        if arr[mid] > arr[right] { // error: arr[mid] > arr[left]  test: a1
            // error: left = mid  test: a
            left = mid + 1	
        } else if arr[mid] < arr[right] { // 属于第二个子数组
			right = mid
		} else {
			right--
		}
	}

	return arr[left]
}

//  err: 死循环
func findMinErr1(arr []int) int {
    left, right := 0, len(arr)-1

    for left < right {
        mid := (left + right) >> 1
        if arr[mid] > right {
            left = mid
        } else if arr[mid] < right {
            right = mid
        } else {
            right--
        }
    }
    return arr[left]
}

func main() {
	a := []int{4, 5, 1, 2, 3}
    a1 := []int{1, 1, 0, 1, 1}
	r := findMin(a1)
    e1 := findMinErr1(a)
	fmt.Println(r)
    fmt.Println(e1)
}
