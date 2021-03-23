package cache

import (
	"time"

	"github.com/go-redis/redis/v8"
)

// Cache interface
type Cache interface {
	Get(key string) interface{}
	Set(key string, val interface{}, timeout time.Duration) error
	IsExist(key string) bool
	Delete(key string) error

	XRead(key string, count int64) []redis.XStream
	XAdd(key, id string, values []string) (string, error)
	XDel(key string, id string) (int64, error)
}
