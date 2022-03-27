package main

import "fmt"

type numeric interface {
	type int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64
}

func max[T numeric](x, y T) T {
	if x > y {
		return x
	}
	return y
}

func min[T numeric](x, y T) T {
	if x > y {
		return y
	}
	return x
}

func main() {
	fmt.Println(min(123, 456))
	fmt.Println(max(123, 456))
}
