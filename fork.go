package main

import "syscall"

func main() {
	ret1, _, err := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
	if err != 0 {
		panic("fork error")
	}
	
}
