package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":7788")
	assert(err, "dial error: %v")
	defer conn.Close()

	datas := []string{
		"data1",
		"data2",
		"data3",
		"data4",
	}

	for i := 0; i < len(datas); i++ {
		io.WriteString(conn, datas[i])
	}
}

func assert(err error, format string) {
	if err != nil {
		panic(fmt.Sprintf(format, err.Error()))
	}
}
