package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func cp(path, to string) {
	read, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("read file error! error: %s", err)
	}
	fmt.Println("please input the copy file name: ")
	in := bufio.NewReader(os.Stdin)
	inRead, _, err := in.ReadLine()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var s = to + string(inRead)
	err = ioutil.WriteFile(s, read, 0666)
	if err != nil {
		fmt.Printf("write file error! error: %s", err)
	}
}

func main() {
	cp("io/t.txt", "io/")
	fmt.Fprintf(os.Stdout, "")
}
