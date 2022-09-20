package main

import (
	"os"
	"path/filepath"
)

func main() {
	home := os.Getenv("HOME")
	fp := filepath.Join(home, "Downloads/jDKUARa.jpg")

	f, err := os.Open(fp)
	if err != nil {
		panic(err)
	}
	defer f.Close()
}
