package main

import (
	"fmt"
	"os"
)

// 1.1 修改 echo 程序输出 os.Args[0]，即命令的名字
func echo() {
	// Output:
	/*
	  /var/folders/b6/spfwtm655mvf64ync5z86s600000gn/T/go-build094758120/b001/exe/p1.1
	 */
	fmt.Println(os.Args[0])
}

func main() {
	echo()
}
