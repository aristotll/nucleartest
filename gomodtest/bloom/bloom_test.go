package bloom

import (
	"sort"
	"testing"
)

func Test1(t *testing.T) {
	bloom1()
}

func Test2(t *testing.T) {
	bloom2()
}

func TestSort(t *testing.T) {
	a := []int{3, 2, 1, 5, 6, 4, 213, 3453, 4313, 123, 43, 5, 135, 6, 456, 1, 456, 56, 576, 53, 45}
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
}
