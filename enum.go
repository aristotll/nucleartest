package main

import (
	"fmt"
)

type IPVersion byte

const (
	IP4 IPVersion = iota + 1
	IP6
)

type OSVersion byte

const (
	LINUX OSVersion = iota + 1
	DARWIN
	WINDOWS
)

func getInfo(v IPVersion) {
	switch v {
	case IP4:
		fmt.Println("IP4")
	case IP6:
		fmt.Println("IP6")
    default:
        fmt.Println("unknown version")
	}
}

func main() {
    fmt.Println(IP4 == LINUX)
	getInfo(0)
	getInfo(IP4)
	//getInfo(LINUX)
}
