package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	name := "73.25"
	h := md5.New()
	h.Write([]byte(name))
	middle := hex.EncodeToString(h.Sum(nil))
	for i := 1; i < 100000000; i++ {
		h.Reset()
		h.Write([]byte(middle))
		middle = hex.EncodeToString(h.Sum(nil))
	}
	fmt.Printf("%s\n", middle)

}
