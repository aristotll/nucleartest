package util

import (
	"encoding/json"
	"log"
)

type Stu struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var s = []Stu{
	{
		Name: "张三",
		Age:  12,
	},
	{
		Name: "张三丰",
		Age:  22,
	},
	{
		Name: "张无忌",
		Age:  33,
	},
}

// 模拟前端返回的 json 数据
func GetJson() []byte {
	data, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
