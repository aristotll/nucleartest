package main

import "fmt"

func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); {
		if nums[i] == nums[nums[i]] {
			return nums[i]
		}
		if nums[i] != i {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		} else if nums[i] == i {
			i++
		}
	}
	return 0
}

func main() {
	n := findRepeatNumber([]int{2, 3, 1, 0, 2, 5, 3})
	fmt.Println(n)
}
