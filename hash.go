package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"hash/crc32"
)

func main() {
	s := "aaa"

	u := crc32.ChecksumIEEE([]byte(s))
	fmt.Printf("u: %x\n", u)

	b := md5.Sum([]byte(s))
	fmt.Printf("b: %x\n", b)

	b2 := sha1.Sum([]byte(s))
	fmt.Printf("b2: %x\n", b2)

	b3 := sha256.Sum256([]byte(s))
	fmt.Printf("b3: %x\n", b3)
}
