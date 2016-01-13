package store

import (
	"testing"
	"time"

	"gopkg.in/alog.v1/log"
)

func TestElasticStore(t *testing.T) {
	var cfg log.ElasticConfig
	cfg.URL = "http://192.168.33.70:9200"
	store := NewElasticStore(cfg)
	var err error
	for i := 0; i < 1000; i++ {
		var item log.LogItem
		item.ID = uint64(i)
		item.Time = time.Now()
		item.Level = log.INFO
		item.Tag = log.DefaultTag
		item.Message = ".........................."
		err = store.Store(&item)
		if err != nil {
			break
		}
	}
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("Write success.")
}
