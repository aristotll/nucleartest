package main

import (
	"encoding/binary"
	"fmt"
)

func test() {
	var localBit uint32 = 0x8002c2f2
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, localBit)
	fmt.Printf("localBit 0x%x to conv netBit is: %v\n", localBit, b)
	v := binary.LittleEndian.Uint32(b)
	fmt.Println(v)

	var bit uint32 = 0x12345678
	binary.LittleEndian.PutUint32(b, bit)
	fmt.Println(b)
	binary.BigEndian.PutUint32(b, bit)
	fmt.Println(b)

	//net.IPv4()
}

func main() {
	test()
}
