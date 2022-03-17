package api

import (
	"context"
	"encoding/json"
	"esDemo/thridpart/model"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var id = 0
var client = model.New()
var index = "student"

// create index
func Create() {
	// create by struct
	s := model.Student{
		Name: "张三",
		Age:  20,
	}

	// client := model.New()
	index0, err := client.Index().
		Index(index).
		Id(strconv.Itoa(id)).
		BodyJson(s).
		Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("put index: %v, id: %v \n", index0.Index, index0.Id)

	// create by json
	j := model.GetJson()
	var sb strings.Builder
	sb.Write(j)

	index1, err := client.Index().
		Index(index).
		Id(strconv.Itoa(id + 1)).
		BodyJson(sb.String()).
		Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("put index: %v, id: %v \n", index1.Index, index1.Id)
}

func Gets() {
	get, err := client.Get().
		Index(index).
		Id(strconv.Itoa(id)).
		Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	if get.Found {
		fmt.Printf("get doc %s in version %d from index %s\n",
			get.Id, get.Version, get.Index)
		// 获取 student 索引下 id 为 0 的文档内容
		js, _ := get.Source.MarshalJSON()
		fmt.Println(string(js))

		var s model.Student
		if err := json.Unmarshal(js, &s); err != nil {
			log.Fatal(err)
		}
		fmt.Println(s)
	}
}

func Update() {
	newData := &model.Student{
		Name: "张三3",
		Age:  10,
	}
	res, err := client.Update().
		Index(index).
		Id(strconv.Itoa(id)).
		Doc(newData).
		Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("update result: %v id: %v \n", res.Result, res.Id)
}

func Search() {

}

func PrintData(result *elastic.SearchResult, err error) {

}
