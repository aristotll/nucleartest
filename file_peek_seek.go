package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	f, err := os.OpenFile("1.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := os.WriteFile("1.txt", []byte("123"), 777); err != nil {
		panic(err)
	}

	b, err := os.ReadFile("1.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

	_, err = f.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	filepath.Glob()

}