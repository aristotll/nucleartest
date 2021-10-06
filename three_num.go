package main

import (
	"fmt"
)

func main() {
	var (
		st, nd, rd int
	)

	var (
		x = 100
		y = 29999
		z = 5
	)

	st = find1st(x, y, z)
	nd = find2nd(st, x, y, z)
	rd = find3rd(st, nd, x, y, z)

	fmt.Println(st, nd, rd)
}

func find1st(x, y, z int) int {
	if x > y && x > z {
		return x
	}
	if y > x && y > z {
		return y
	}
	return z
}

func find2nd(st, x, y, z int) int {
	switch st {
	case x:
		return max(y, z)
	case y:
		return max(x, z)
	case z:
		return max(x, y)
	default:
		return 0
	}
}

func find3rd(st, nd, x, y, z int) int {
	if x != st && x != nd {
		return x
	} else if y != st && y != nd {
		return y
	}
	return z
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
