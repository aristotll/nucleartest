package main

import "log"

var L = &Logs{l: &log.Logger{}}

type Logs struct {
	l *log.Logger
}

func main() {
	L.l.Println("123")
}
