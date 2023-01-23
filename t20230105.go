package main

import "fmt"

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

	fmt.Println(err1 == err2, err2 == err3, err1 == err2)
}
