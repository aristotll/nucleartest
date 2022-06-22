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

func Test2(t *testing.T) {
	fmt.Printf("%x\n", 0x7fffffffe820-0x8)
	fmt.Printf("%x\n", 0x7fffffffe818-0x8)
	fmt.Printf("%x\n", 0x820-0x8)
	fmt.Printf("%x\n", 0x818-0x8)
}
