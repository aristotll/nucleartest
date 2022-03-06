package es

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

var EsClient *elasticsearch.Client

func init() {
	var err error
	EsClient, err = elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s \n", err)
	}

	resp, err := EsClient.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer resp.Body.Close()
	if resp.IsError() {
		log.Fatalf("Error: %s", resp.String())
	}

	var r map[string]interface{}
	// Deserialize the response into a map.
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print client and server version numbers.
	log.Printf("Client: %s", elasticsearch.Version)
	log.Printf("Server: %s", r["version"].(map[string]interface{})["number"])
	log.Println(strings.Repeat("~", 37))
}

func Put(es *elasticsearch.Client, req *esapi.IndexRequest) error {
	resp, err := req.Do(context.TODO(), es)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.IsError() {
		log.Printf("[%s] Error indexing document ID=%v", resp.String(), req.DocumentID)
	} else {
		// Deserialize the response into a map.
		var r map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and indexed document version.
			log.Printf("[%s] %s; version=%d",
				resp.Status(), r["result"], int(r["_version"].(float64)))
		}
	}
	return nil
}
