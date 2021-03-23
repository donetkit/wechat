package cache

import (
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	opts := &RedisOpts{
		Host:     "127.0.0.1:6379",
		Password: "",
		Database: 0,
	}
	redis := NewRedis(opts)
	var err error
	timeoutDuration := 100 * time.Second

	if err = redis.Set("username", "TestRedis", timeoutDuration); err != nil {
		t.Error("set Error", err)
	}

	if !redis.IsExist("username") {
		t.Error("IsExist Error")
	}

	name := redis.Get("username").(string)
	if name != "TestRedis" {
		t.Error("get Error")
	}

	if err = redis.Delete("username"); err != nil {
		t.Errorf("delete Error , err=%v", err)
	}
}
