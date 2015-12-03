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
	lists  *list.List
	locker sync.Mutex
}

func (mb *_MemoryBuffer) Push(item log.LogItem) error {
	mb.locker.Lock()
	defer mb.locker.Unlock()
	mb.lists.PushBack(item)
	return nil
}

func (mb *_MemoryBuffer) Pop() (*log.LogItem, error) {
	mb.locker.Lock()
	defer mb.locker.Unlock()
	ele := mb.lists.Front()
	if ele == nil || ele.Value == nil {
		return nil, nil
	}
	item := ele.Value.(log.LogItem)
	mb.lists.Remove(ele)
	return &item, nil
}
