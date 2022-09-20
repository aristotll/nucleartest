package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type I interface {
	Do()
}

type _S struct {
	//pri  int64
	Id   int64
	//Name []byte
}

func (s *_S) Do() { fmt.Println("Do") }

var _ I = &_S{}

func write(b io.Writer, i I) {
	if err := binary.Write(b, binary.BigEndian, i); err != nil {
		panic(err)
	}
}

func read(b io.Reader, i I) {
	if err := binary.Read(b, binary.BigEndian, i); err != nil {
		panic(err)
	}
}

func main() {
	var (
		buf = new(bytes.Buffer)
		//_buf = make([]byte, 4096)
		//buf  = bytes.NewBuffer(_buf) 这样写会导致无法 read 出数据，原因未知
		s = &_S{
			Id:   10001,
			//Name: []byte{'a', 'b', 'c'},
		}
		s_ = &_S{}
	)
	write(buf, s)
	read(buf, s_)
	fmt.Println(s_)
	s_.Do()
}
