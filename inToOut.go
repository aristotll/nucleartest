package main

import (
	"fmt"
	"io"
	"os"
)

// 计算标准输入的数字总和
func CountInSum() int {
	sum, curVal := 0, 0
	for {
		_, err := fmt.Scan(&curVal)
		if err != nil {
			fmt.Println("input error !")
			break
		}
		sum += curVal
	}
	return sum
}

// 标准输入读到标准输出
func InToOut() {
	io.Copy(os.Stdout, os.Stdin)
}

func main() {
	// io.Copy(os.Stdout, os.Stdin)
	fmt.Println(CountInSum())
}
