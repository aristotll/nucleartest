package main

import (
	"os"
	"fmt"
)

func main() {
	curDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	
	fmt.Println(curDir)
}
