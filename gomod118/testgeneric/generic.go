package testgeneric

import (
	"constraints"
	"fmt"
)

func Max[T constraints.Ordered](x, y T) T {
	if x > y {
		return x
	}
	return y
}

func Slices[T comparable](sli []T) {
	fmt.Println(sli)
}

type Struct[T any] struct {
	A, B T
}

func NewStruct[T any](a, b T) *Struct[T] {
	return &Struct[T]{
		A: a,
		B: b,
	}
}

func (s *Struct[T]) Print() {
	fmt.Println(s)
}

func ReturnT[T any](t T) T {
	return t
}
