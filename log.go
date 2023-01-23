package main

import (
	"log"
	"os"
)

//var L = &Logs{l: &log.Logger{}}	// 空指针
var L = &Logs{l: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)}

type Logs struct {
	l *log.Logger
}

func main() {
	L.l.Println("123")
}
