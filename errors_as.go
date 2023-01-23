package main

import (
	"errors"
	"fmt"
)

var _ error = &QueryError{}

type QueryError struct {
	Msg       string
	QueryPath string
}

func (e *QueryError) Error() string {
	return fmt.Sprintf("query %s error: %s", e.QueryPath, e.Msg)
}

func main() {
	var err1 = &QueryError{
		Msg:       "network error",
		QueryPath: "/a/b/c",
	}
	var err2 = &QueryError{
		Msg:       "unknown error",
		QueryPath: "a/b/c",
	}
	var err3 = &QueryError{
		Msg:       "network error",
		QueryPath: "/a/b/c",
	}
	var err *QueryError

	// fmt.Println(err1 == err2, err2 == err3)
	// errors.Is 会判断两个 error 是否是同一个对象，只有是同一个对象的情况下才会返回 true
	// 即便两个 error 是同一种类型
	fmt.Printf("errors.Is(err1, err2): %v\n", errors.Is(err1, err2))
	fmt.Printf("errors.Is(err2, err3): %v\n", errors.Is(err2, err3))
	fmt.Printf("errors.Is(err1, err): %v\n", errors.Is(err1, err))

	// errors.As 只会判断两个 error 是否是同一种类型，只要是同一种类型就会返回 true
	// 比如下面的 err1 和 err 都是 *QueryError 类型的，As 就会返回 true
	// 此外 As 还会做一项额外的工作：如果两者为同一类型，将 err（第一个参数）赋值
	// 给 target（第二个参数）
	// 注意 As 的第二个参数必须加 &，即便已经是指针类型
	fmt.Println(errors.As(err1, &err))
	// As 会将 err1 的错误信息写入到 err 中，所以这里输出 network error 
    fmt.Println(err.Error())

	fmt.Printf("errors.As(err2, err1): %v\n", errors.As(err2, &err1))
	fmt.Println(err1.Error())

	var Err = errors.New("err")
	var Err1 *QueryError
	fmt.Printf("errors.As(Err, &Err1): %v\n", errors.As(Err, &Err1))
	fmt.Printf("Err1: %v\n", Err1)
}
