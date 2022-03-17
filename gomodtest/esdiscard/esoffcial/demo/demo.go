package demo

import (
	"context"
	"encoding/json"
	"esDemo/esoffcial/util"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"strconv"
	"strings"
)

// 使用
// elasticsearch 包将两个单独的包绑定在一起，分别用于调用 elasticsearch api
// 和通过 HTTP 传输数据：esapi 和 estransport。
// 使用elasticsearch.NewDefaultClient（）函数创建具有默认设置的客户端。
func Usage() {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("esoffcial --version: ", elasticsearch.Version)
	fmt.Println(es.Info())

	rsp, err := es.Info()
	if err != nil {
		log.Fatal(err)
	}

	// 注意：为了在默认的HTTP传输中重用持久的TCP连接，关闭响应体和使用响应体是非常重要的。
	// 如果您对响应主体不感兴趣，请调用io.copy(ioutil.Discard, res.Body).
	defer rsp.Body.Close()
	log.Println(rsp)
}

func API() {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatal(err)
	}

	jsonData := util.GetJson()
	// fmt.Printf("%s\n", jsonData)

	var sb strings.Builder
	id := 1

	sb.Write(jsonData)
	// 设置请求对象
	req := esapi.IndexRequest{
		Index:      "student",
		DocumentID: strconv.Itoa(id),
		Body:       strings.NewReader(sb.String()),
		Refresh:    "true",
	}

	rsp, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer rsp.Body.Close()
	if rsp.IsError() {
		log.Printf("[%s] Error indexing document ID=%d", rsp.Status(), id)
	} else {
		// Deserialize the response into a map.
		var r map[string]interface{}

		if err := json.NewDecoder(rsp.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and indexed document version.
			log.Printf("[%s] %s; version=%d", rsp.Status(), r["result"], int(r["_version"].(float64)))
		}
	}
}
