package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		fmt.Println("create error: ", err)
		return
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		fmt.Println("trace start error: ", err)
		return
	}
	defer trace.Stop()

	fmt.Println("Hello World")
}
