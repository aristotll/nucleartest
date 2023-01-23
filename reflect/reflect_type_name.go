package main

import (
	"fmt"
	"reflect"
)

type S struct{}

type I interface{}

func pointer[T any](v T) *T {
	return &v
}

func GetTypeName(t any) string {
	typ := reflect.TypeOf(t)
	if typ.Kind() != reflect.Pointer {
		panic("must pointer")
	}
	typ = typ.Elem()
	return typ.Name()
}

func printF(t any, fn func(t any) string) {
	fmt.Println(fn(t))
}

func main() {
	s := new(S)
	i32 := pointer(123)
	i64 := pointer(456)
	str := pointer("abc")
    var itf I

	printF(s, GetTypeName)
	printF(i32, GetTypeName)
	printF(i64, GetTypeName)
	printF(str, GetTypeName)
    printF(itf, GetTypeName)
}
