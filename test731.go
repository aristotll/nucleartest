package main

import (
	
)

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	var res int

	for i < j {
		curval := min(height[i], height[j]) * (j - i)
		res = min(res)
	}
}
