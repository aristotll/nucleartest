package main

import (
	"fmt"
)

type TT struct {
	name string
}

func tslice() []TT {
	n := make([]*TT, 1, 1) // make([]*TT, 1, 1) does not escape
	t := &TT{}             // &TT{} escapes to heap
	n[0] = t

	nn := make([]TT, 1, 1) // make([]TT, 1, 1) escapes to heap
	t1 := TT{}
	nn[0] = t1

	a := make([]int, 1024) // make([]int, 1024) does not escape
	a[0] = 1

	a1 := make([]int, 4096) // make([]int, 4096) does not escape
	a1[0] = 1

	a2 := make([]int, 10000) // make([]int, 10000) escapes to heap
	a2[0] = 1

	return nn
}

func tnew() {
	t := new(TT) // new(TT) does not escape
	_ = t
}

func main() {
	n := tslice()
	fmt.Println(n) // []interface {}{...} does not escape
	tnew()
}
