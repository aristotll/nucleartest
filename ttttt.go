package main

import (
	"fmt"
	"reflect"
)

// ps -ef | grep ttttt | grep -v grep | grep run | awk {'print $2'} | xargs kill

type I interface {
	Do()
}

type S struct{}

func (s *S) Do() {}

func EnforcePtr(obj interface{}) {
	v := reflect.TypeOf(obj)
    fmt.Println(v)
	if v == nil {
        typ := reflect.TypeOf(&v).Elem()
		fmt.Println(typ.Kind())
	}
}

func main() {
	var i I
	EnforcePtr(i)

    var n S
    EnforcePtr(n)

	var ii I = new(S)
	EnforcePtr(ii)
}
