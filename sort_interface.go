package main

import (
	"fmt"
	"sort"
)

type canSort []int

func (c canSort) Len() int {
	return len(c)
}

func (c canSort) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c canSort) Less(i, j int) bool {
	return c[i] < c[j]
}

type stu1 struct {
	name string
	age int8
}

type ss []*stu1



func main() {
	var c canSort
	c = append(c, 1)
	c = append(c, -10)
	c = append(c, -0)
	c = append(c, 346)

	sort.Sort(c)

	fmt.Println(c)
}
