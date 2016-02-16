package buffer

import (
	"testing"
	"time"

	"github.com/antlinker/alog/log"
)

func TestRedisPush(t *testing.T) {
	config := log.RedisConfig{
		Addr: "192.168.33.70:6379",
		DB:   1,
	}
	redisBuf := NewRedisBuffer(config)
	var item log.LogItem
	item.ID = 1
	item.Time = time.Now()
	item.Tag = log.DefaultTag
	item.Level = log.DEBUG
	item.Message = "Test redis message..."
	err := redisBuf.Push(item)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("Success")
}

func TestRedisFront(t *testing.T) {
	config := log.RedisConfig{
		Addr: "192.168.33.70:6379",
		DB:   1,
	}
	redisBuf := NewRedisBuffer(config)
	item, err := redisBuf.Front()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(item)
}

func TestRedisPop(t *testing.T) {
	config := log.RedisConfig{
		Addr: "192.168.33.70:6379",
		DB:   1,
	}
	redisBuf := NewRedisBuffer(config)
	item, err := redisBuf.Pop()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(item)
}
