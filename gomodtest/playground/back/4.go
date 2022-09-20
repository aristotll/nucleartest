package main

import "fmt"

func rmP(s *[]int) {
	*s = (*s)[:len(*s)-1]
	//fmt.Println(s)
}

func rm(s []int) {
	s = s[:len(s)-1]
}

func main() {
	s := []int{1, 2, 4, 5, 99}

	rm(s)
	fmt.Println(s)

	rmP(&s)
	fmt.Println(s)
}
