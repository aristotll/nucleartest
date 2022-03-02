package main

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	fmt.Printf("%d \n", 0xe)
	fmt.Printf("%d \n", 0xf)
	fmt.Printf("%d \n", 0x13)
	fmt.Println(0xc == 0x13)
}

func Test11(t *testing.T) {
	fmt.Println((1 + 9) >> 1)
}
