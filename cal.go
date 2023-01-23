package main

import "fmt"

func main() {
    fmt.Printf("0xfe = %b\n", 0xfe)
	fmt.Printf("(11101010 & 0xfe) = %b\n", (11101010 & 0xfe))
	fmt.Printf("(11101010 & 11111110) = %b\n", (11101010 & 11111110))
    fmt.Printf("0000 0011 & 0000 0101 = %b\n", 00000011 & 00000101)
}
