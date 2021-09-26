package main

import (
	"fmt"
	"reflect"
	"testing"
)

type Cul struct {
	X, Y int64
}

func (c *Cul) Add() int64 {
	return c.X + c.Y
}


func TestServer(t *testing.T) {
	c := new(Cul)
	c.X = 10
	c.Y = 20

	fn := reflect.ValueOf(c).MethodByName("Add")
	res := fn.Call([]reflect.Value{})
	
	for _, r := range res {
		fmt.Println(r.Interface())
	}
}
