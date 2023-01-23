package main

import (
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"fmt"
	"hash"
)

type ss struct {
	hashFunc hash.Hash
}

func new(fn hash.Hash) *ss {
	s := &ss{}
	if fn == nil {
		s.hashFunc = md5.New()
	} else {
		s.hashFunc = fn
	}
	return s
}

func (s *ss) set(v any) {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(v); err != nil {
		panic(err)
	}
	s.hashFunc.Write(buf.Bytes())
	fmt.Println(buf.Bytes())
}

func main() {
	type st struct {
		Name string
		Age  int8
	}
	s := new(nil)
	s.set(&st{"abc", 18})
	s.set(123)
	s.set("abc")
}
