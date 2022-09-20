package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type DB struct {
	M map[string]string
	A []int
}

func encode() {
	d := &DB{
		M: map[string]string{
			"a": "1",
			"b": "2",
			"c": "3",
		},
		A: []int{1, 2, 3},
	}

	f, err := os.OpenFile("gob_struct.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	enc := gob.NewEncoder(f)
	if err := enc.Encode(d); err != nil {
		panic(err)
	}
}

func decode() {
	f, err := os.Open("gob_struct.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var d DB
	dec := gob.NewDecoder(f)
	if err := dec.Decode(&d); err != nil {
		panic(err)
	}
	if v, ok := d.M["a"]; ok {
		fmt.Printf("a:%v\n", v)
	}
	fmt.Println(d.A)
}

func main() {
	encode()
	decode()
}
