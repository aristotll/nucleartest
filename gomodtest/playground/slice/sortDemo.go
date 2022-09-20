package main

import (
	"fmt"
	"sort"
)

type stu struct {
	name string
	age int
}

func sortD(s *[]stu) {
	fmt.Printf("排序前 ：%+v \n", s)
	// fmt.Println((*s)[0].age)
	sort.SliceStable(*s, func(i, j int) bool {
		if (*s)[i].age < (*s)[j].age {
			return true
		}
		return false
	})
	fmt.Printf("排序后：%+v \n", s)
}

func main() {
	m := make([]stu, 5)
	m[0] = stu{
		name: "zhang3",
		age:  16,
	}
	m[1] = stu{
		name: "li4",
		age:  66,
	}
	m[2] = stu{
		name: "li4",
		age:  3,
	}
	sortD(&m)
}
