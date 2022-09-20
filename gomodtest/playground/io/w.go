package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
)

func wt() {
	file, err := os.OpenFile("t.txt", os.O_RDWR|os.O_APPEND, 0744)
	if err != nil {
		log.Fatal(err)
	}
	n, err := file.Write([]byte("\n test123 by writer \n"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("writer %d bytes \n", n)

	w := bufio.NewWriter(file)
	nn, err := w.Write([]byte("writer by bufio \n "))
	if err != nil {
		log.Fatal(err)
	}
	err = w.Flush()
	fmt.Printf("bufio writer %d bytes \n", nn)
}

func main() {
	//wt()
	fmt.Println(10 << 2)
	fmt.Println(10 >> 2)
	d := 1129.6
	fmt.Println(d * 100)
	v := big.NewFloat(d * 100)
	
	f, _ := v.Float64()
	fmt.Println(f)
}
