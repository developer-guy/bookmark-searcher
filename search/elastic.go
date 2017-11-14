package search

import (
	"gopkg.in/olivere/elastic.v5"
	"fmt"
)

const SCHEME = "http"

type Options []elastic.ClientOptionFunc

type ElasticOpt struct {
	ElasticHost string
}

func prepareCloudConfigs(client *ElasticOpt) Options {
	var options Options = Options{}
	options = append(options, elastic.SetURL(client.ElasticHost))
	options = append(options, elastic.SetScheme(SCHEME))
	options = append(options, elastic.SetSniff(false))
	return options
}

func (elasticClient *ElasticOpt) Client() (*elastic.Client, error) {
	var options Options = prepareCloudConfigs(elasticClient)
	client, err := elastic.NewSimpleClient(options...)

	if err != nil {
		return nil, err
	}

	fmt.Println("Elastic client initialized successfully...")
	return client, nil
}
