package main

import (
	"fmt"
	"math/big"
)

func main() {
	var bits big.Int
	for i := 1000; i < 2000; i++ {
		bits.SetBit(&bits, i, 1)
	}
	for i := 0; i < 10000; i++ {
		if bits.Bit(i) != 0 {
			fmt.Println(i)
		}
	}
}
