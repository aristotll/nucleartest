package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(get2(9, 32))
}

func get2(ones, bits int) string {
	count1 := ones
	count0 := bits - ones

	var has8 int8
	var ret strings.Builder
	for i := 0; i < count1; i++ {
		if has8 == 8 {
			ret.WriteString(".")
			has8 = 0
		}
		ret.WriteString("1")
		has8++
	}

	for i := 0; i < count0; i++ {
		if has8 == 8 {
			ret.WriteString(".")
			has8 = 0
		}
		ret.WriteString("0")
		has8++
	}

	return ret.String()
}
