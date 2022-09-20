package es

import (
	"context"
	"testing"

	"github.com/olivere/elastic/v7"
)

var EsClientOlivere *elastic.Client

func init() {
	cli, err := elastic.NewClient()
	if err != nil {
		panic(err)
	}
	EsClientOlivere = cli
}

type Employee struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Age       int16    `json:"age"`
	About     string   `json:"about"`
	Interests []string `json:"interests"`
}

func TestPutOlivere(t *testing.T) {
	req := &Employee{
		FirstName: "John",
		LastName:  "Smith",
		Age:       25,
		About:     "I love to go rock climbing",
		Interests: []string{"sports", "music"},
	}
	// mapping 类似 mysql 中的表结构定义
	mapping := `
	{
		"mappings":{
		   "employee":{
			  "properties":{
				 "first_name":{
					"type":"keyword"
				 },
				 "last_name":{
					"type":"keyword"
				 },
				 "age":{
					"type":"integer"
				 },
				 "about":{
					"type":"text"
				 },
				 "interests":{
					"type":"keyword"
				 }
			  }
		   }
		}
	 }`
	// es7 不再支持定义 type，所以需要使用下面的写法
	mapping1 := `
	{
		"settings": {
			"analysis": {
			  "normalizer": {
				"lowercase": {
				  "type": "custom",
				  "filter": ["lowercase"]
				}
			  }
			}
		},
		"mappings":{
			  "properties":{
				 "first_name":{
					"type":"keyword",
					"normalizer": "lowercase"
				 },
				 "last_name":{
					"type":"keyword",
					"normalizer": "lowercase"
				 },
				 "age":{
					"type":"integer"
				 },
				 "about":{
					"type":"text"
				 },
				 "interests":{
					"type":"keyword"
				 }
			  }
		}
	 }`
	_ = mapping
	indexName := "megacorp"
	//id := "1"
	if err := PutOlivere(context.Background(), EsClientOlivere, mapping1, indexName, "1", req); err != nil {
		t.Fatal(err)
	}

	req1 := &Employee{
		FirstName: "Douglas",
		LastName:  "Fir",
		Age:       35,
		About:     "I like to build cabinets",
		Interests: []string{"forestry"},
	}
	if err := PutOlivere(context.Background(), EsClientOlivere, mapping1, indexName, "2", req1); err != nil {
		t.Fatal(err)
	}

	req2 := &Employee{
		FirstName: "Jane",
		LastName:  "Smith",
		Age:       32,
		About:     "I like to collect rock albums",
		Interests: []string{"music"},
	}
	if err := PutOlivere(context.Background(), EsClientOlivere, mapping1, indexName, "3", req2); err != nil {
		t.Fatal(err)
	}
}
func TestSearchOlivere(t *testing.T) {
	indexName := "megacorp"
	id := "1"
	ctx := context.Background()
	var e Employee
	if err := SearchOlivere(ctx, EsClientOlivere, indexName, id, &e); err != nil {
		t.Fatal(err)
	}
	t.Logf("search result: %+v\n", e)
}

func TestSearchDSL(t *testing.T) {
	indexName := "megacorp"
	// 	GET /megacorp/_search
	// 	{
	//     "query" : {
	//         "bool": {
	//             "must": {
	//                 "match" : {
	//                     "last_name" : "smith"
	//                 }
	//             },
	//             "filter": {
	//                 "range" : {
	//                     "age" : { "gt" : 30 }
	//                 }
	//             }
	//         }
	//     }
	// }
	// 下面的 dsl 语句等同于上面的 json，也可以调用 bq.Source() 来查看
	bq := elastic.NewBoolQuery()
	bq = bq.Must(elastic.NewTermQuery("last_name", "smith"))
	bq = bq.Filter(elastic.NewRangeQuery("age").Gt(30))
	var emp Employee

	res, err := SearchDSL(context.Background(), EsClientOlivere, bq, indexName, &emp)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range res {
		t.Logf("%+v\n", v)
	}

	// 在`about` 属性上搜索 “rock climbing”
	mq := elastic.NewMatchQuery("about", "rock climbing")

	// Elasticsearch 默认按照相关性得分排序，即每个文档跟查询的匹配程度。
	// 第一个最高得分的结果很明显：John Smith 的 about 属性清楚地写着 “rock climbing” 。
	// 但为什么 Jane Smith 也作为结果返回了呢？原因是她的 about 属性里提到了 “rock” 。
	// 因为只有 “rock” 而没有 “climbing” ，所以她的相关性得分低于 John 的。
	// 这是一个很好的案例，阐明了 Elasticsearch 如何 在 全文属性上搜索并返回相关性最强的结果。
	// Elasticsearch中的 相关性 概念非常重要，也是完全区别于传统关系型数据库的一个概念，
	// 数据库中的一条记录要么匹配要么不匹配。
	res, err = SearchDSL(context.Background(), EsClientOlivere, mq, indexName, &emp)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range res {
		t.Logf("%+v\n", v)
	}

	// match_phrase 用于精确匹配一系列单词或者_短语_ 。 比如，想查询仅匹配同时
	// 包含 “rock” 和 “climbing” ，并且二者以短语 “rock climbing” 的形式紧
	// 挨着的雇员记录。
	mpq := elastic.NewMatchPhraseQuery("about", "rock climbing")
	res, err = SearchDSL(context.Background(), EsClientOlivere, mpq, indexName, &emp)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range res {
		t.Logf("%+v\n", v)
	}
}

func TestSearchHighlight(t *testing.T) {
	mq := elastic.NewMatchQuery("about", "rock climbing")

	h := elastic.NewHighlight()
	h = h.Fields(elastic.NewHighlighterField("about")).
		PreTags("<span style='color: red;'>").
		PostTags("</span>")

	indexName := "megacorp"

	r, err := SearchHighlight(context.Background(), EsClientOlivere, mq, h, indexName)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range r {
		t.Logf("%+v\n", v.Highlight)
	}
}
