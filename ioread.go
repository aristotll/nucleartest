package main

import (
	"io"
	"log"
	"strings"
)

func main() {
	var (
		buf  = make([]byte, 10)
		str  = strings.NewReader("12345")
		str1 = strings.NewReader("abcde")
	)

	_, err := io.ReadFull(str, buf[:2])
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(buf)

	_, err = io.ReadFull(str1, buf[:2])
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(buf)
}
