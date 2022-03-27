package main

import (
	"fmt"
	"github.com/samber/lo"
	"strconv"
)

func main() {
	fmt.Println(lo.Union([]int{1, 2, 3, 4}, []int{1, 10, 3, 4}))
	fmt.Println(lo.Uniq([]int{1, 2, 3, 3}))
	fmt.Println(lo.Uniq([]string{"abc", "abc", "def", "def"}))
	fmt.Println(lo.Max([]int{2, 1, -99, 30, 50, 3, 99, 1}))
	fmt.Println(lo.Map[int64, string]([]int64{1, 2, 3, 4, 5}, func(x int64, _ int) string {
		return strconv.FormatInt(x, 10)
	}))

}
