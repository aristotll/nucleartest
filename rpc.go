package main

import (
	"fmt"
	"reflect"
)

type St struct{}

func main() {
	s := &St{}
	styp := reflect.TypeOf(s)
	fmt.Println(styp.String())

	sval := reflect.ValueOf(s)
	name := reflect.Indirect(sval).Type().Name()
	fmt.Println(name)

	ss := &struct{}{}
	ssval := reflect.ValueOf(ss)
	name = reflect.Indirect(ssval).Type().Name()
	fmt.Println(name)
}
