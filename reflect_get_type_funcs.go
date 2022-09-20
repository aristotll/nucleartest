package main

import (
	"fmt"
	"reflect"
)

type Struct struct{}

func (s *Struct) X()               {}
func (s *Struct) Y(x, y int) error { return nil }

func main() {
	typ := reflect.TypeOf(new(Struct))
	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		mtyp := method.Type

		fmt.Println(mtyp.NumIn())
		arg1 := mtyp.In(0)
		fmt.Printf("kind: %v, name: %v \n", arg1.Kind(), arg1.String())
	}
}
