package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := "1@2:3"
	sep := "@"
	before, after, found := bytes.Cut([]byte(s), []byte(sep))
	fmt.Printf("Cut(%q, %q) = %q, %q, %v\n", s, sep, before, after, found)
}
