package main

import (
	"flag"
	"fmt"
	"strconv"
)

var input = flag.String("n", "", "输入一个数字，二进制以 0b 开头，八进制以 0o 开头，十六进制以 0x 开头，十进制直接输入数字")

func main() {
	flag.Parse()
	bit := (*input)[:2]
	if bit[0] == '0' {
		switch bit[1] {
		case 'b':
			getRes((*input)[2:], 2)
		case 'o':
			getRes((*input)[2:], 8)
		case 'x':
			getRes((*input)[2:], 16)
		default:
			getRes(*input, 10)
		}
	} else {
		getRes(*input, 10)
	}
}

func getRes(v string, bit int8) {
	i10, err := strconv.ParseInt(v, int(bit), 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("10 进制：%v\n", i10)
	bs := []int{2, 8, 16}
	for _, b := range bs {
        s := strconv.FormatInt(i10, b)
		fmt.Printf("%v 进制：%v\n", b, s)
	}
}
