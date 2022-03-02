package reflect

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	var i int64 = 123
	//var ii = &i

	var newVal reflect.Value
	typ := reflect.TypeOf(i)
	argIsValue := false
	//t.Log(typ.String())
	//t.Log(typ.Elem())
	if typ.Kind() == reflect.Pointer {
		t.Log("typ is a point")
		newVal = reflect.New(typ.Elem())
	} else {
		t.Log("typ is a value")
		newVal = reflect.New(typ)
		argIsValue = true
	}
	if argIsValue {
		newVal = newVal.Elem()
	}
	t.Log(newVal.Interface())
	t.Log(newVal.Elem())
}

func Test2(t *testing.T) {
	var i int64
	typ := reflect.TypeOf(i)
	newi := reflect.New(typ)
	t.Log(newi.Type().String())
	t.Log(newi.Interface())
	t.Log(newi.Elem().Interface())
	t.Log("===============")

	var j *int64 = &i
	typ = reflect.TypeOf(j)
	newi = reflect.New(typ)
	t.Log(newi.Type().String())
	t.Log(newi.Interface())
	t.Log(newi.Elem().Interface())
}

type Struct struct{}

func (s *Struct) XX(x, y int64) int64 {
	return x + y
}

// call method
func Test3(t *testing.T) {
	s := &Struct{}
	method := reflect.TypeOf(s).Method(0)
	p1typ := method.Type.In(0) // s
	p2typ := method.Type.In(1) // x
	p3typ := method.Type.In(2) // y

	p1 := reflect.New(p1typ.Elem())
	p2 := reflect.New(p2typ).Elem()
	p3 := reflect.New(p3typ).Elem()

	var a int64 = 2
	var b int64 = 3
	p2.Set(reflect.ValueOf(a))
	p3.Set(reflect.ValueOf(b))

	returnValues := method.Func.Call([]reflect.Value{p1, p2, p3})

	retv := returnValues[0].Interface()
	t.Log(retv)
}
