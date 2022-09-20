package main

import (
	"fmt"
	"io"
	"os"
)

// 输入一个 io.Write，返回一个 io.Write

type Obj struct {
	msg string
}

func (o *Obj) Write(p []byte) (n int, err error) {
	*o = Obj{
		msg: string(p),
	}
	return len(o.msg), nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	// 传递的参数是 os.Stdin，所以下面的 w.Write() 即 os.Stdin.Write()
	// os.Stdin.Write(): 打印到终端，与 fmt.Print() 相同
	w.Write([]byte("sadsadad\n"))
	o := Obj{}
	return &o, nil
}

func main() {
	writer, i := CountingWriter(os.Stdin)
	n, _ := writer.Write([]byte("123"))
	fmt.Println(writer)
	fmt.Println(i)
	fmt.Println(n)
}
