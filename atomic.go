package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// a,b 两个线程并发写入到同一个文件，不使用 os.O_APPEND ，使用 seek 移动到末尾

func FileSize(f *os.File) int {
	b := make([]byte, 1024)
	//var b [1024]byte
	var size int
	for {
		n, err := f.Read(b)
		if err != nil && err != io.EOF {
			fmt.Println("get file size error: ", err)
			return -1
		}
		if err == io.EOF {
			break
		}
		size += n
	}
	return size
}

func AppendFile() {
	file, err := os.OpenFile("./testgeneric.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		log.Fatal("open file error: ", err)
	}

	n := FileSize(file)
	seek, _ := file.Seek(int64(n), io.SeekEnd)
	fmt.Println("seek: ", seek)
	_, err = file.Write([]byte("123abc\n"))
	if err != nil {
		fmt.Println("write to file error: ", err)
		return
	}
	fmt.Println("执行了 1 次")
	// sta , _ := file.Stat()
	// log.Println("file size: ", n)
	// log.Println("file size [stat]", sta.Size())
}

func main() {
	for i := 0; i < 3; i++ {
		go AppendFile()
		go AppendFile()
		go AppendFile()
	}
	time.Sleep(time.Second * 5)
}
