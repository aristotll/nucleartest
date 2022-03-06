package es

import (
	"context"
	"encoding/json"
	"log"
	"reflect"

	"github.com/olivere/elastic/v7"
)

// 参考：https://strconv.com/posts/use-elastic/

// elasticsearch7 默认不再支持指定索引类型，默认索引类型是 _doc
// 查询：curl -X GET "localhost:9200/[indexname]/_doc/[id]"
func PutOlivere(
	ctx context.Context,
	c *elastic.Client,
	mapping, indexName, id string,
	body interface{},
) error {
	// 检查 index 是否已经存在，不存在则创建
	exists, err := c.IndexExists(indexName).Do(ctx)
	if err != nil {
		return err
	}
	if !exists {
		_, err = c.CreateIndex(indexName).BodyString(mapping).Do(ctx)
		if err != nil {
			log.Printf("create index error: %v", err)
			return err
		}
	}

	resp, err := c.Index().
		Index(indexName).
		//Type(typeName).
		Id(id).
		BodyJson(body).
		Refresh("wait_for").
		Do(ctx)
	if err != nil {
		log.Printf("put error: %v", err)
		return err
	}
	log.Printf("Indexed with id=%v, type=%s\n", resp.Id, resp.Type)
	return nil
}

func SearchOlivere(
	ctx context.Context,
	c *elastic.Client,
	indexName, id string,
	struct_ interface{},
) error {
	result, err := c.Get().Index(indexName).Id(id).Do(ctx)
	if err != nil {
		return err
	}
	if result.Found {
		log.Printf("Got document %v (version=%d, index=%s, type=%s)\n",
			result.Id, result.Version, result.Index, result.Type)
		if err := json.Unmarshal(result.Source, struct_); err != nil {
			return err
		}
	}
	return nil
}

func SearchDSL(
	ctx context.Context,
	cli *elastic.Client,
	query elastic.Query,
	indexName string,
	struct_ interface{},
) (res []interface{}, err error) {
	resp, err := cli.Search().Index(indexName).Query(query).Do(ctx)
	if err != nil {
		return nil, err
	}
	res = resp.Each(reflect.TypeOf(struct_))
	return
}

func SearchHighlight(
	ctx context.Context,
	cli *elastic.Client,
	query elastic.Query,
	high *elastic.Highlight,
	indexName string,
) (r []*elastic.SearchHit, err error) {
	res, err := cli.Search(indexName).Highlight(high).Query(query).Do(ctx)
	if err != nil {
		return nil, err
	}
	r = res.Hits.Hits
	return
}
