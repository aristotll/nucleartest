package main

type A[T any] struct {
	a T
}

func NewA[T any](a T) *A[T] {
	return &A[T]{
		a: a,
	}
}

type B struct {
	a A[int]
}

func main() {
	var b B

}
