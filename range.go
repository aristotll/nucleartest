package main

import (
	"fmt"
)

type Struct struct {
	X int64
}

func main() {
	//type Struct struct {X int64}
	list := []*Struct{{1}, {2}, {3}}
	cp := make([]*Struct, len(list))

	for k, v := range list {
		cp[k] = v
	}
	printSlice(list, "list")
	printSlice(cp, "cp")
	//fmt.Println("list: ", list)
	//fmt.Println("cp: ", cp)

	list1 := []Struct{{1}, {2}, {3}}
	cp1 := make([]*Struct, len(list))
	for k, v := range list1 {
		cp1[k] = &v
	}
	printSlice1(list1, "list1")
	printSlice(cp1, "cp1")
}

func printSlice(sli []*Struct, name string) {
	fmt.Printf("%v: ", name)
	for _, v := range sli {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}

func printSlice1(sli []Struct, name string) {
	fmt.Printf("%v: ", name)
	for _, v := range sli {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}
