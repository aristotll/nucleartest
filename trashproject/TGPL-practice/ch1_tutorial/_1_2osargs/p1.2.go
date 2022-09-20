package main

import (
	"fmt"
	"os"
)

// 1.2 修改 echo 程序，输出参数的索引和值，每行一个
func echo1() {
	for k, v := range os.Args {
		fmt.Println(k, v)
	}
}

func main() {
	// go run p1.2.go param1 param2 param3
	// Output:
	// 0 /var/folders/b6/spfwtm655mvf64ync5z86s600000gn/T/go-build876781442/b001/exe/p1.2
	// 1 param1
	// 2 param2
	// 3 param3
	echo1()
}
