package main

import (
	"fmt"
	"os"
)

func main() {
	curDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fmt.Println(curDir)
}
