package main

import (
	"fmt"
	"go/token"
	"reflect"
)

type i64 int64

type St struct{}
type st struct{}

func main() {
	var (
		a int64
		b = &struct{ x, y int64 }{}
		c i64
		d = &St{}
		e = &st{}
	)
	println(isExportedOrBuiltinType(reflect.TypeOf(a))) // true
	println(isExportedOrBuiltinType(reflect.TypeOf(b))) // true
	println(isExportedOrBuiltinType(reflect.TypeOf(c))) // false
	println(isExportedOrBuiltinType(reflect.TypeOf(d))) // true
	println(isExportedOrBuiltinType(reflect.TypeOf(e))) // false

	at := reflect.TypeOf(a)
	i64t := reflect.TypeOf(c)
	fmt.Println(at.PkgPath(), i64t.PkgPath())
}

// Is this type exported or a builtin?
func isExportedOrBuiltinType(t reflect.Type) bool {
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	// PkgPath will be non-empty even for an exported type,
	// so we need to check the type name as well.
	return token.IsExported(t.Name()) || t.PkgPath() == ""
}
