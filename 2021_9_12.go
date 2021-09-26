package main

import (
	"reflect"
)

func add(x, y int64) int64 {
	return x + y
}

func main() {
	f := add
	fn := reflect.ValueOf(f)
	param1 := reflect.ValueOf(int64(1))
	param2 := reflect.ValueOf(int64(2))

	call := fn.Call([]reflect.Value{param1, param2})

	_ = call
}
