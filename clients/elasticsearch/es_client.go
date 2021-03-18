package elasticsearch

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/agusluques/bookstore_utils-go/logger"
	"github.com/olivere/elastic"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(*elastic.Client)
	Index(string, string, interface{}) (*elastic.IndexResponse, error)
}

type esClient struct {
	client *elastic.Client
}

func Init() {
	log := logger.GetLogger()
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetErrorLog(log),
		elastic.SetInfoLog(log),
	)
	if err != nil {
		panic(err)
	}
	Client.setClient(client)
}

func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}

func (c *esClient) Index(index string, typ string, doc interface{}) (*elastic.IndexResponse, error) {
	result, err := c.client.Index().
		Index(index).
		Type(typ).
		BodyJson(doc).
		Do(context.Background())

	if err != nil {
		log.Fatal(fmt.Sprintf("error when trying to index document in index %s.", index), err)
	}

	return result, err
}
