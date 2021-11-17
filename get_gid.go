package main

import (
    "fmt"
    "runtime"
    "strings"
    "strconv"
)

func goID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

func main() {
    for i := 0; i < 100; i++ {
        go func(i int) {
           fmt.Printf("go %d ID: %d \n", i, goID())    
        }(i)
    }

    g := goID()
    fmt.Println(g)

    select{}
}
