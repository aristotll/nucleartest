package es

import (
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func TestPut(t *testing.T) {
	json_ := `
	{
		"first_name" : "John",
		"last_name" :  "Smith",
		"age" :        25,
		"about" :      "I love to go rock climbing",
		"interests": [ "sports", "music" ]
	}
	`
	// b, err := json.Marshal(json_)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	err := Put(EsClient, &esapi.IndexRequest{
		Index:      "megacorp/employee/1",
		DocumentID: "1",
		Body:       strings.NewReader(json_),
		Refresh:    "true",
	})
	if err != nil {
		t.Fatal(err)
	}
}
