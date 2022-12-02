package database

import (
	"bytes"
	"context"
	"fmt"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

func LoadDataToElasticsearch(source_id string, data []byte) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		// ...
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	request := esapi.IndexRequest{Index: source_id, DocumentID:"1", Body: bytes.NewReader(data)}

	res, err := request.Do(context.Background(), es)
	defer res.Body.Close()
	if res.IsError() {
		fmt.Printf("[%s] Error indexing document %s", res.Status(), res)
	}

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return
}
