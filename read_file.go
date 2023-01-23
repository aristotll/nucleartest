package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.OpenFile("testfile.txt", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		panic(err)
	}

	size := fi.Size()
	fmt.Printf("size: %v\n", size)
	if size == 0 {
		f.WriteString("123")
	}

	b := make([]byte, 20)
	f.Seek(0, 0)
	_, err = f.Read(b)
	if err != nil {
		panic(err)
	}

	fmt.Println(b)
}
