package main

import (
	"runtime"
	"fmt"
)

func main() {
	sysType := runtime.GOOS
	fmt.Println(sysType)
}
