package main

import (
	"fmt"
	"io"
	"time"
)

func main() {
	r, w := io.Pipe()
	go func() {
		n, err := w.Write([]byte("123"))
		if err != nil {
			fmt.Println("write error: ", err)
			return
		}
		fmt.Printf("write %d bytes\n", n)
	}()

	buf := make([]byte, 100)

	go func() {
		n, err := r.Read(buf)
		if err != nil {
			fmt.Println("read error: ", err)
			return
		}
		fmt.Printf("read %d bytes\n", n)
		fmt.Println(string(buf))
	}()

	time.Sleep(time.Second * 3)
}
