package main

import (
	"os"
	"fmt"
)

func main() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	fmt.Println("os.Executable: ", ex)

	ex1, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("os.Getwd: ", ex1)
}
