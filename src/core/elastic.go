package core

import (
	"context"
	"log"
	"os"

	"github.com/olivere/elastic"
)

var ES *elastic.Client

func ElasticConnect() {
	var err error
	ES, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(os.Getenv("elastic_host")))
	if err != nil {
		//panic(err)
	}
}

func CloseElasticConnect() {
	ES.Stop()
}

func getElasticInit() {
	index := os.Getenv("log_index")
	exist, err := ES.IndexExists(index).Do(context.Background())
	if err != nil {
		log.Printf("check is index exist, got: %v", err)
	}
	if exist {
		return
	}
	createIndex, err := ES.CreateIndex(index).Do(context.Background())
	if err != nil {
		log.Printf("try to create index, got: %v", err)
	}
	if createIndex == nil {
		log.Printf("try to create index and expect != nil, got: %v", createIndex)
	}
	mapping := ``
	_, err = ES.PutMapping().Index(index).BodyString(mapping).Do(context.TODO())
	if err != nil {
		log.Printf("put mapping got: %v", err)
	}

	getMapping, err := ES.GetMapping().Index(index).Do(context.TODO())
	if err != nil {
		log.Printf("expected get mapping to succeed; got: %v", err)
	}
	_, ok := getMapping[index]
	if !ok {
		log.Printf("expected JSON root to be of type map[string]interface{}; got: %#v", ok)
	}

	return
}

func ElasticInsert(index, data string) bool {
	getElasticInit()
	_, err := ES.Index().Index(index).Type("_doc").BodyString(data).Do(context.Background())

	if err != nil {
		log.Printf("expect insert into elasticsearch got: %v", err)
		return false
	}
	return true
}
