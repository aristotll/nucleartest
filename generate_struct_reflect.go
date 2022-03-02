package main

import (
	"errors"
	"fmt"
	"reflect"
)

type s struct {
	// 构造器
	// 用于存储属性字段
	fileId []reflect.StructField
}

func NewBuilder() *Builder {
	return &Builder{}
}

// 添加字段，传入字段名和类型
func (b *Builder) AddField(field string, typ reflect.Type) *Builder {
	b.fileId = append(b.fileId, reflect.StructField{Name: field, Type: typ})
	return b
}

// 根据预先添加的字段构建出结构体类型
func (b *Builder) Build() *Struct {
	// StructOf 返回包含字段的 struct 类型
	stu := reflect.StructOf(b.fileId)
	index := make(map[string]int)
	for i := 0; i < stu.NumField(); i++ {
		index[stu.Field(i).Name] = i
	}
	return &Struct{
		typ:   stu,
		index: index,
	}
}

func (b *Builder) AddString(name string) *Builder {
	return b.AddField(name, reflect.TypeOf(""))
}

func (b *Builder) AddBool(name string) *Builder {
	return b.AddField(name, reflect.TypeOf(true))
}

func (b *Builder) AddInt64(name string) *Builder {
	return b.AddField(name, reflect.TypeOf(int64(0)))
}

func (b *Builder) AddFloat64(name string) *Builder {
	return b.AddField(name, reflect.TypeOf(float64(1.2)))
}

func (b *Builder) AddStringSlice(name string) *Builder {
	return b.AddField(name, reflect.TypeOf([]string{}))
}

// 实际生成的结构体，基类
// 结构体的类型
type Struct struct {
	typ reflect.Type
	// <fieldName : 索引> // 用于通过字段名称，从Builder的[]reflect.StructField中获取reflect.StructField
	index map[string]int
}

// 根据结构体类型创建一个实例化
func (s Struct) New() *Instance {
	return &Instance{
		instance: reflect.New(s.typ).Elem(),
		index:    s.index,
	}
}

// 结构体的值
type Instance struct {
	instance reflect.Value
	// <fieldName : 索引>，赋值时需要根据字段名定位到其在结构体中的索引，
	// 再通过索引定位到该字段地址，进而进行赋值操作
	index map[string]int
}

var (
	FieldNoExist error = errors.New("field no exist")
)

func (in Instance) Field(name string) (reflect.Value, error) {
	if i, ok := in.index[name]; ok {
		return in.instance.Field(i), nil
	} else {
		return reflect.Value{}, FieldNoExist
	}
}

func (in *Instance) SetString(name, value string) {
	if i, ok := in.index[name]; ok {
		in.instance.Field(i).SetString(value)
	}
}

func (in *Instance) SetBool(name string, value bool) {
	if i, ok := in.index[name]; ok {
		in.instance.Field(i).SetBool(value)
	}
}

func (in *Instance) SetInt64(name string, value int64) {
	if i, ok := in.index[name]; ok {
		in.instance.Field(i).SetInt(value)
	}
}

func (in *Instance) SetFloat64(name string, value float64) {
	if i, ok := in.index[name]; ok {
		in.instance.Field(i).SetFloat(value)
	}
}

func (in *Instance) SetStringSlice(name string, value []string) {
	if i, ok := in.index[name]; ok {
		v := reflect.ValueOf(value)
		in.instance.Field(i).Set(v)
	}
}

func (i *Instance) Interface() interface{} {
	return i.instance.Interface()
}

func (i *Instance) Addr() interface{} {
	return i.instance.Addr().Interface()
}

func main() {
	pe := NewBuilder().
		AddString("Name").
		AddInt64("Age").
		AddFloat64("Height").
		AddStringSlice("Scores").
		Build()

	p := pe.New()
	p.SetString("Name", "你好")
	p.SetInt64("Age", 32)
	p.SetFloat64("Height", 100)
	p.SetStringSlice("Scores", []string{"1", "2"})

	fmt.Printf("%+v\n", p)
	fmt.Printf("%T，%+v\n", p.Interface(), p.Interface())
	fmt.Printf("%T，%+v\n", p.Addr(), p.Addr())
}
