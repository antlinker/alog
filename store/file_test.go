package store

import (
	"testing"
	"time"

	"github.com/antlinker/alog/log"
)

func TestFileStore(t *testing.T) {
	var config log.FileConfig
	config.FileSize = 10
	store := NewFileStore(config)
	var err error
	for i := 0; i < 1000; i++ {
		var item log.LogItem
		item.ID = uint64(i)
		item.Time = time.Now()
		item.Level = log.DEBUG
		item.Tag = log.DefaultTag
		item.Message = "............."
		item.File.Name = "file_test.go"
		item.File.Line = 22
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
