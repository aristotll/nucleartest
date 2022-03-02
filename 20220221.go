package main

import "fmt"

func moveZeroes(nums []int) {
	//  i  j
	// [0, 1, 0, 3, 12]
	// n[i] == 0 && n[j] != 0;swap(n[i], n[j]);i++
	//
	//     i  j
	// [1, 0, 0, 3, 12]
	// [1, 3, 0, 0, 12]
	//
	//
	i := 0
	j := 0
	for i < len(nums) {
		j = i
		// j 移动到非 0 处
		for j < len(nums) && nums[j] == 0 {
			j++
		}

		if nums[i] == 0 && j < len(nums) {
			nums[i], nums[j] = nums[j], nums[i]
		}

		i++
		//fmt.Println(i, j, nums)
	}
}

func main() {
	n := []int{0, 1, 0, 3, 12}
	moveZeroes(n)
	fmt.Println(n)
}
