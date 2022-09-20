package main

import "fmt"

func c1(s *[]int, i int) {
	//f := (*s)[i:]
	//f1 := (*s)[i+1:]
	//fmt.Println(f)
	//fmt.Println(f1)
	copy((*s)[i:], (*s)[i+1:])
	*s = (*s)[:len(*s)-1]
}

func rm(s *[]int, i int) {
	// [12 34 54]
	front := (*s)[:i]
	// [90]
	last := (*s)[i+1:]

	for _, v := range last {
		front = append(front, v)
	}
	*s = front
}

func rm2(s *[]int, i int) {
	//// [12 34 54]
	//front := (*s)[:i]
	//// [90]
	//last := (*s)[i+1:]
	//
	//*s = append(front, last...)
	*s = append((*s)[:i], (*s)[i+1:]...)
}

func main() {
	s := make([]int, 0)
	s = append(s, 12, 34, 54, 78, 90)
	//rm(&s, 2)
	//fmt.Println(s)

	//c1(&s, 2)
	//fmt.Println(s)

	rm2(&s, 2)
	fmt.Println(s)
}
