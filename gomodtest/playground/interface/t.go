package main

import (
	"fmt"
	"io"
)

// io.Write 接口的实现演示

type ByteCounter int
type Model1 string

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return int(*c), nil
}

func (m *Model1) Write(p []byte) (int, error) {
	*m = Model1(p)
	return 0, nil
}

func aWriter(w io.Writer, s string) {
	_, _ = w.Write([]byte(s))
	// fmt.Println(write)
}

func main() {
	var b ByteCounter
	aWriter(&b, "123")
	fmt.Println(b)

	var m Model1
	aWriter(&m, "sadaf")
	fmt.Println(m)

}
