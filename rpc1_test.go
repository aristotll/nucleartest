package main

import (
	"log"
	"net/http"
	"net/rpc"
	"testing"
)

type Ss struct {
	X, Y int64
	Str string
}

func (s *Ss) Add(param *Param, res *Res) error {
	res.Add = param.X + param.Y
	res.Str = param.Str
	return nil
}


type Param struct {
	Str string
	X, Y int64
}

type Res struct {
	Str string
	Add int64
}

// 成员顺序和 Param 不一样，如果 client 使用这个结构体作为参数， rpc 能否调用成功？
// 结论：可以
type ParamDif struct {
	X, Y int64
	Str string
}

// 少了一个字段，能否调用成功？
// 结论：可以
type ParamDif1 struct {
	X, Y int64
}

// 字段名不同，能否调用成功？
// 失败：gob: type mismatch: no fields matched compiling decoder for Param
type ParamDif2 struct {
	XXX, YYY int64
}

// 多一个字段
// 结论：成功
type ParamDif3 struct {
	Str string
	X, Y int64
	III float64
}

func TestServer1(t *testing.T) {
	if err := rpc.Register(new(Ss)); err != nil {
		log.Fatalln(err)
	}
	rpc.HandleHTTP()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}

func TestClient1(t *testing.T) {
	client, err := rpc.DialHTTP("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	par := new(ParamDif3)
	par.X = 10
	par.Y = 20
	par.Str = "aaa"
	par.III = 123.1

	res := new(Res)

	if err := client.Call("Ss.Add", par, res); err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}