package store

import (
	"text/template"

	"gopkg.in/alog.v1/log"

	"gopkg.in/mgo.v2"
)

// NewMongoStore 创建基于MongoDB存储的实例
func NewMongoStore(cfg log.MongoConfig) log.LogStore {
	if cfg.URL == "" {
		cfg.URL = log.DefaultMongoURL
	}
	if cfg.DBTmpl == "" {
		cfg.DBTmpl = log.DefaultMongoDBTmpl
	}
	if cfg.CollectionTmpl == "" {
		cfg.CollectionTmpl = log.DefaultMongoCollectionTmpl
	}
	session, err := mgo.Dial(cfg.URL)
	if err != nil {
		panic(err)
	}
	return &MongoStore{
		session:        session,
		dbTmpl:         template.Must(template.New("").Parse(cfg.DBTmpl)),
		collectionTmpl: template.Must(template.New("").Parse(cfg.CollectionTmpl)),
	}
}

type MongoStore struct {
	session        *mgo.Session
	dbTmpl         *template.Template
	collectionTmpl *template.Template
}

func (ms *MongoStore) Store(item *log.LogItem) error {
	dbName := log.ParseName(ms.dbTmpl, item)
	collectionName := log.ParseName(ms.collectionTmpl, item)
	err := ms.session.DB(dbName).C(collectionName).Insert(item.ToMap())
	return err
}
