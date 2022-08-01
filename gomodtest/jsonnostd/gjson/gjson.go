package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"jsonnostd"
	"strconv"
)

// 获取单个字段
func getSingleValue() {
	res := gjson.Get(jsonnostd.GlobalJsonString, "data.date_time").String()
	fmt.Println(res)
}

// 从数组中通过下标获取
func getFromArrayIndex() {
	str := gjson.Get(jsonnostd.GlobalJsonString, "data.search_data*.0").String()
	result := gjson.Get(str, "*.0")
	PrintMap(result.Map())
	PrintMap(result.Map()["location"].Map())
}

func getArray() {
	var array []gjson.Result
	g := gjson.Get(jsonnostd.GlobalJsonString, "data.elements")
	if g.IsArray() {
		array = gjson.Get(jsonnostd.GlobalJsonString, "data.elements").Array()
	} else {
		panic("not array")
	}

	var dataArray []gjson.Result
	dataRes := array[4].Map()["data"]
	if dataRes.IsArray() {
		dataArray = dataRes.Array()
	}
	PrintMap(dataArray[0].Map())
}

func getArrayLast() {
	var array []gjson.Result
	g := gjson.Get(jsonnostd.GlobalJsonString, "data.elements")
	if g.IsArray() {
		array = gjson.Get(jsonnostd.GlobalJsonString, "data.elements").Array()
	} else {
		panic("not array")
	}
	PrintMap(g.Array()[len(g.Array())-1].Map())

	size := len(array) - 1
	sizeStr := strconv.Itoa(size)
	path := "data.elements." + sizeStr
	fmt.Println(path)
	PrintMap(gjson.Get(jsonnostd.GlobalJsonString, path).Map())
}

func getArraySize() {
	g := gjson.Get(jsonnostd.GlobalJsonString, "data.elements.#")
	fmt.Println(g.String())
}

func main() {
	//getArray()
	//getArrayLast()
	getArraySize()
}

func PrintMap(m map[string]gjson.Result) {
	for k, v := range m {
		fmt.Printf("%v:%v\n", k, v)
	}
}
