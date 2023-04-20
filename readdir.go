package main

import (
	"fmt"
	"os"
)

func main() {
	dirs, err := os.ReadDir("./helm")
	if err != nil {
		panic(err)
	}
	for _, dir := range dirs {
		fmt.Println(dir.Name())
        _, err := os.ReadFile(dir.Name())
        if err != nil {
            panic(err)
        }
	}
}
