package main

import (
	"fmt"
	"github.com/docker/docker/pkg/reexec"
	"log"
	"os"
	"syscall"
	"time"
)

var mmap, globalErr = syscall.Mmap(-1, 0, 4096,
	syscall.PROT_READ|syscall.PROT_WRITE,
	syscall.MAP_SHARED|syscall.MAP_ANON)

func init() {
	if globalErr != nil {
		panic(globalErr)
	}
	reexec.Register("child", childProcessFunc)
	if reexec.Init() {
		os.Exit(1)
	}
}

func childProcessFunc() {
	time.Sleep(time.Second)
	fmt.Printf("child got a message: %s\n", mmap)
	copy(mmap, "hi, dad, this is son")
	fmt.Println("child write ok")
	os.Exit(0)
}

func main() {
	cmd := reexec.Command("child")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		log.Panicf("failed to run command: %s", err)
	}
	//if err := cmd.Wait(); err != nil {
	//	log.Panicf("failed to wait command: %s", err)
	//}
	copy(mmap, "hi, dad, this is father")
	fmt.Println("father write ok")
	time.Sleep(time.Second * 3)
	fmt.Printf("father got a message: %s\n", mmap)
	log.Println("main exit")
}
