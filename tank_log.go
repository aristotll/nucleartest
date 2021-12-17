package main

import (
	"fmt"
)

func Log(format string, v ...interface{}) {
	content := fmt.Sprintf(format+"\r\n", v...)
	fmt.Println(content)
}

func main() {
	Log("a: %v \n", 1)
	Log("detect packet %v failed: %v. ", 1, 2, 3)
	Log("a")
}
