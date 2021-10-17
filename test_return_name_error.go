package main

import (
	"encoding/binary"
	"fmt"
)

func retErr() error {
	return fmt.Errorf("error!")
}

func fn() (err error) {
	if err = retErr(); err != nil {
		err = fmt.Errorf("aaa")
	}
	return
}

func main() {
	if err := fn(); err != nil {
		fmt.Println(err)
	}
	fmt.Println([]byte("8080"))
	fmt.Println([]byte{80, 80})
	fmt.Println([]byte{8, 0, 8, 0})

	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, 8080)
	fmt.Println(b)

	u := binary.BigEndian.Uint16(b)
	fmt.Println(u)
}