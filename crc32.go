package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	s := "abcdefg"
	h := crc32.ChecksumIEEE([]byte(s))
	fmt.Println(h)
}
