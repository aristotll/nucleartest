package main

import "fmt"
import "sort"

/**
 * 删除重复元素
 * @param array int整型一维数组
 * @return int整型一维数组
 */
func removeDuplicate(array []int) []int {
	// write code here
	m := make(map[int]int)
	for i, v := range array {
		m[v] = i
	}
	a := make([]int, len(array), len(array))
	copy(a, array)
	for i := 0; i < len(array); i++ {
		if lastIndex, ok := m[array[i]]; ok {
			a = append(a[:lastIndex], a[lastIndex+1:]...)
		}
		m[array[i]] = i
	}
	return a
}

func removeDuplicate1(array []int) []int {
	// write code here
	m := make(map[int]int)
	var a []int
	copy(a, array)
	for i := len(array) - 1; i >= 0; i-- {
		if index, ok := m[array[i]]; ok {
			a = append(a[:index], a[index+1:]...)
		}
		m[array[i]] = i
	}
	return a
}

func main() {
	//ret := removeDuplicate([]int{3, 5, 8, 2, 3, 8})
	//fmt.Println(ret)

	ret1 := removeDuplicate1([]int{3, 5, 8, 2, 3, 8})
	fmt.Println(ret1)
}
