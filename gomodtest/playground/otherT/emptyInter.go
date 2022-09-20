package main

import "fmt"

type T interface{}


type stu struct {
	name string
	age  int
}

// param: 空接口类型
func t(param T) T {
	return param
}

// 空接口类型的转换
// p: 接收一个空接口参数，并传参给 f()
// f: 一个参数为空接口，返回值为空接口的函数
// 该非常恶心的函数的作用：将 【f】 的【空接口返回值】 转换为【具体的类型】
func getType(p T, f func(T) T) {
	v := f(p)
	// 为了将一个接口变量转化为一个显式的类型，又可以用 .(TYPE)
	// 如果底层类型不是 int，将输出 err
	// i := v.(int)

	// 或者使用类型转换
	switch v.(type) {
	case int:
		fmt.Println("参入的参数为 int 类型, value: ", v)
	case string:
		fmt.Println("参入的参数为 string 类型, value: ", v)
	case *stu:
		fmt.Println("参入的参数为 *stu 类型, value: ", v)
	}
}

func main() {
	s := stu{
		name: "zzz",
		age:  123,
	}
	getType(&s, func(i T) T {
		return i
	})
}
