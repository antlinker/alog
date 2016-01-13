package buffer

import (
	"container/list"
	"sync"

	"gopkg.in/alog.v1/log"
)

// NewMemoryBuffer 创建新的Memory实例
func NewMemoryBuffer() log.LogBuffer {
	return &_MemoryBuffer{
		lists: list.New(),
	}
}

// _MemoryBuffer 内存缓冲区
type _MemoryBuffer struct {
	sync.RWMutex
	lists *list.List
}

func (mb *_MemoryBuffer) Push(item log.LogItem) error {
	mb.Lock()
	mb.lists.PushBack(item)
	mb.Unlock()
	return nil
}

func (mb *_MemoryBuffer) Pop() (*log.LogItem, error) {
	mb.Lock()
	ele := mb.lists.Front()
	if ele == nil || ele.Value == nil {
		mb.Unlock()
		return nil, nil
	}
	v := ele.Value
	mb.lists.Remove(ele)
	mb.Unlock()
	item := v.(log.LogItem)
	return &item, nil
}
