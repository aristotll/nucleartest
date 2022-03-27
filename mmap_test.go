package main

import (
	"fmt"
	"os"
	. "syscall"
	"testing"
	"time"
)

// 使用 mmap 进行进程间通信
// go test -run ^TestRead$ -v mmap_test.go
// go test -run ^TestWrite$ -v mmap_test.go

const defaultMmapPageSize = 1024

// 该进程负载写入
func TestWrite(t *testing.T) {
	fmt.Printf("process id=%v, this process will write data\n", os.Getpid())
	f, err := os.OpenFile("./mmap.txt", os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data, err := Mmap(int(f.Fd()), 0, defaultMmapPageSize, PROT_READ|PROT_WRITE, MAP_SHARED)
	if err != nil {
		panic(err)
	}
	// 写入数据
	str := "test test"
	for i := 0; i < len(str); i++ {
		data[i] = str[i]
	}
}

func TestRead(t *testing.T) {
	fmt.Printf("process id=%v, this process will read data\n", os.Getpid())
	f, err := os.OpenFile("./mmap.txt", os.O_RDWR, 0777)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	data, err := Mmap(int(f.Fd()), 0, defaultMmapPageSize, PROT_READ|PROT_WRITE, MAP_SHARED)
	if err != nil {
		t.Fatal(err)
	}

	for {
		fmt.Printf("data: %v\n", string(data))
		time.Sleep(time.Second * 2)
	}
}

func TestAnonymousMmap(t *testing.T) {
	data, err := Mmap(-1, 0, defaultMmapPageSize, PROT_READ|PROT_WRITE, MAP_SHARED|MAP_ANON)
	if err != nil {
		t.Fatal(err)
	}

	ret1, ret2, err := Syscall(SYS_FORK, 0, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	
}
