package main

import (
	"fmt"
	"reflect"
)

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{3, 4, 5}
	r := reflect.DeepEqual(s1, s2)
	fmt.Println(r)
	r = reflect.DeepEqual(s2, s3)
	fmt.Println(r)
}
