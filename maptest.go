package main

import (
    "fmt"
)

type MessageType interface {
	TypeName() string
}

var globalMsgID int64

type xxx struct {
	name string
}

func (x *xxx) TypeName() string {
	return ""
}

func GenMsgTyp(typName string) MessageType {
	return &xxx{name: typName}
}

func main() {
    m := make(map[MessageType]struct{})
    m[GenMsgTyp("http")] = struct{}{}
    m[GenMsgTyp("rpc")] = struct{}{}
    fmt.Println(m)   
}
