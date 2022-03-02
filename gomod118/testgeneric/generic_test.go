package testgeneric

import "testing"

func TestMax(t *testing.T) {
	//println(Max(1, 2))
}

func TestSlices(t *testing.T) {
	Slices([]string{"1", "2"})

	type S struct {
		A, B int64
	}

	Slices([]S{
		{1, 2},
		{3, 4},
	})

	type SS[T any] struct {
		A, B T
	}

	Slices([]SS[int64]{
		{1, 2},
		{3, 2},
	})

	// error: []int does not implement comparable
	//Slices([][]int{
	//	{1, 2, 3},
	//	{4, 5, 6},
	//})
}

func TestStruct_Print(t *testing.T) {
	NewStruct(1, 2).Print()
}

func TestReturnT(t *testing.T) {
	println(ReturnT(1))
}
