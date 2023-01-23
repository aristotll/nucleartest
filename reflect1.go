package main

import (
	"fmt"
	"reflect"
)

type I interface {
	Do()
}

type S struct{}

func (s *S) Do() {}

func EnforcePtr(obj interface{}) {
	v := reflect.ValueOf(obj)
	fmt.Println(v.Kind())
}

func EnforcePtr1(obj interface{}) {
	v := reflect.TypeOf(obj)
	if v == nil {
		fmt.Println(v.Kind())
	}
}

func main() {
	var i I
	EnforcePtr(i)
	EnforcePtr1(i)

	var ii I = new(S)
	EnforcePtr(ii)
	EnforcePtr1(ii)
}
