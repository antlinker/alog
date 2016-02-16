package buffer

import (
	"testing"
	"time"

	"github.com/antlinker/alog/log"
)

func TestMemoryPush(t *testing.T) {
	memBuf := NewMemoryBuffer()
	var item log.LogItem
	item.ID = 1
	item.Time = time.Now()
	item.Tag = log.DefaultTag
	item.Level = log.DEBUG
	item.Message = "Test message..."
	memBuf.Push(item)
	lItem, _ := memBuf.Pop()
	t.Log(lItem)
	lItem, _ = memBuf.Front()
	t.Log(lItem)
}

func TestMemoryFront(t *testing.T) {
	memBuf := NewMemoryBuffer()
	lItem, err := memBuf.Front()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(lItem)
}

func TestMemoryPop(t *testing.T) {
	memBuf := NewMemoryBuffer()
	lItem, err := memBuf.Pop()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(lItem)
}
