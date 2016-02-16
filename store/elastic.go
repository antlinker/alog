package store

import (
	"text/template"

	"github.com/antlinker/alog/log"
	"gopkg.in/olivere/elastic.v3"
)

// NewElasticStore 创建基于ElasticSearch存储的实例
func NewElasticStore(cfg log.ElasticConfig) log.LogStore {
	if cfg.URL == "" {
		cfg.URL = log.DefaultElasticURL
	}
	if cfg.IndexTmpl == "" {
		cfg.IndexTmpl = log.DefaultElasticIndexTmpl
	}
	if cfg.TypeTmpl == "" {
		cfg.TypeTmpl = log.DefaultElasticTypeTmpl
	}
	client, err := elastic.NewClient(elastic.SetURL(cfg.URL), elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return &ElasticStore{
		client:    client,
		indexTmpl: template.Must(template.New("").Parse(cfg.IndexTmpl)),
		typeTmpl:  template.Must(template.New("").Parse(cfg.TypeTmpl)),
	}
}

// ElasticStore 提供ElasticSearch的存储实现
type ElasticStore struct {
	client    *elastic.Client
	indexTmpl *template.Template
	typeTmpl  *template.Template
}

func (es *ElasticStore) Store(item *log.LogItem) error {
	indexName := log.ParseName(es.indexTmpl, item)
	typeName := log.ParseName(es.typeTmpl, item)
	_, err := es.client.Index().Index(indexName).Type(typeName).BodyJson(item).Do()
	return err
}
