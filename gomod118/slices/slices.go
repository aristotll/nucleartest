package main

import (
	"fmt"
	"golang.org/x/exp/slices"
)

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{6, 7, 8, 9, 10}

	s11 := slices.Delete(s1, 1, 3)
	fmt.Println(s11)

	index := slices.Index(s1, 3)
	println(index)

	fmt.Println(slices.Equal(s1, s2))
	fmt.Println(slices.Equal(s1, []int{1, 2, 3, 4, 5}))

	s3 := make([]int, 0, 10)
	s3 = append(s3, 1, 2, 3)
	fmt.Println(s3, cap(s3))
	s3 = slices.Clip(s3)
	fmt.Println(s3, cap(s3))

	ss3 := slices.Clone(s3)
	fmt.Println(ss3)

	ss33 := slices.Insert(s3, 1, 5)
	fmt.Println(ss33)

	slices.Sort(ss33)
	fmt.Println(ss33)
}
