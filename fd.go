package main

import (
	"os"
)

func main() {
	f, err := os.Open("")
	if err != nil {
		panic(err)
	}
	f.Fd()
}
