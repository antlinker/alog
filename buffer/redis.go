package buffer

import (
	"encoding/json"
	"fmt"

	"gopkg.in/alog.v1/log"
	"gopkg.in/redis.v3"
)

const (
	_LISTKEY = "ALOGS"
)

// NewRedisBuffer 创建新的Redis实例
func NewRedisBuffer(config log.RedisConfig) log.LogBuffer {
	opts := new(redis.Options)
	if v := config.Network; v != "" {
		opts.Network = v
	}
	if v := config.Addr; v != "" {
		opts.Addr = v
	}
	if v := config.PoolSize; v > 0 {
		opts.PoolSize = v
	}
	opts.DB = config.DB
	client := redis.NewClient(opts)
	return &_RedisBuffer{
		wRlient: client.Pipeline(),
		rClient: client,
		dClient: client,
	}
}

type _RedisBuffer struct {
	wRlient *redis.Pipeline
	rClient *redis.Client
	dClient *redis.Client
}

func (rb *_RedisBuffer) Push(item log.LogItem) error {
	key := fmt.Sprintf("%d_%d", item.Time.Unix(), item.ID)
	val := rb.encode(item)
	rb.wRlient.Set(key, val, 0)
	// if err := itemResult.Err(); err != nil {
	// 	return err
	// }
	rb.wRlient.RPush(_LISTKEY, key)
	// if err := listResult.Err(); err != nil {
	// 	return err
	// }
	_, err := rb.wRlient.Exec()
	return err
}

func (rb *_RedisBuffer) Pop() (*log.LogItem, error) {
	result := rb.rClient.LPop(_LISTKEY)
	if result.Err() == redis.Nil {
		return nil, nil
	}
	itemKey, err := result.Result()
	if err != nil {
		return nil, err
	}
	result = rb.rClient.Get(itemKey)
	if result.Err() == redis.Nil {
		return nil, nil
	}
	itemData, err := result.Bytes()
	if err != nil {
		return nil, err
	}
	if len(itemData) == 0 {
		return nil, nil
	}
	item := rb.decode(itemData)
	rb.delKey(itemKey)
	return item, nil
}

func (rb *_RedisBuffer) encode(item log.LogItem) []byte {
	buf, _ := json.Marshal(item)
	return buf
}

func (rb *_RedisBuffer) decode(data []byte) *log.LogItem {
	var item log.LogItem
	json.Unmarshal(data, &item)
	return &item
}

func (rb *_RedisBuffer) delKey(key string) error {
	result := rb.dClient.Del(key)
	return result.Err()
}
