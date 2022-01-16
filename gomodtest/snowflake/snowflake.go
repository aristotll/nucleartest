package main

import (
	"fmt"
	"github.com/sony/sonyflake"
)

func main() {
	sf := sonyflake.NewSonyflake(sonyflake.Settings{})

	for i := 0; i < 100; i++ {
		fmt.Println(sf.NextID())
	}

}
