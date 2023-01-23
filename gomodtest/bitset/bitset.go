package main

import (
	"fmt"

	"github.com/bits-and-blooms/bitset"
)

func main() {
	bs := bitset.New(10086)
	bs.Set(1).Set(2)
	bs.Set(1).Set(2)
	fmt.Printf("bs.Count(): %v\n", bs.Count())
	fmt.Printf("bs.Test(1): %v, bs.Test(10): %v\n", bs.Test(1), bs.Test(10))
	bs.Clear(1)
	fmt.Printf("count: %v, test(1): %v\n", bs.Count(), bs.Test(1))

	bs.Set(1).Set(2)
	u, b := bs.NextSet(0)	// 从下标 0 开始，找到第一个设置为 1 的位
	fmt.Printf("idx: %v, find: %v\n", u, b)	// idx: 1, find: true

	bs.ClearAll()
	u, b = bs.NextSet(0)	// 从下标 0 开始，找到第一个设置为 1 的位
	fmt.Printf("idx: %v, find: %v\n", u, b)	// idx: 0, find: false

	bs.Set(0).Set(2)
	u2, b2 := bs.NextClear(0)	// 从 0 开始
	fmt.Printf("idx: %v, find: %v\n", u2, b2)

	u2, b2 = bs.NextClear(0)	// 从 0 开始
	fmt.Printf("idx: %v, find: %v\n", u2, b2)
}