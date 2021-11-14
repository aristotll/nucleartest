package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	src := []byte("abc123")
	l := hex.DecodedLen(len(src))
	b := make([]byte, l) 

	_, err := hex.Decode(b, src)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("b: %v\n", b)
}