package model

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
)

var client *elastic.Client
var host = "http://127.0.0.1:9200/"

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func New() *elastic.Client {
	client, err := elastic.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	// ping
	info, code, err := client.Ping(host).Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n",
		code, info.Version.Number)

	version, err := client.ElasticsearchVersion(host)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("es version is: %v \n", version)
	return client
}

func GetJson() []byte {
	j, err := json.Marshal(Student{
		Name: "张三丰",
		Age:  30,
	})
	if err != nil {
		log.Fatal(err)
	}
	return j
}
