package main

import (
	"fmt"
)

type size int64

const (
	b       = iota
	kb size = 1 << (10 * iota)
	mb
	gb
)

func main() {
	fmt.Println(b, kb, mb, gb)
}
