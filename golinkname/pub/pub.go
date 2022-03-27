package pub

import (
	_ "golinkname/private"
	_ "unsafe"
)

// private.privateFunc 链接到这里
func callPrivateFunc(string)

// 链接到 private.privateFunc1
//go:linkname callPrivateFunc1 golinkname/private.privateFunc1
func callPrivateFunc1(string, int8)

func CallPrivateFunc() {
	callPrivateFunc("zhang3")
}

func CallPrivateFunc1() {
	callPrivateFunc1("li4", 18)
}

