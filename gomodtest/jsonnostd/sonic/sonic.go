package main

import (
	"github.com/bytedance/sonic"
	"jsonnostd"
)

func main() {
	sonic.UnmarshalString(jsonnostd.GlobalJsonString, nil)
}
