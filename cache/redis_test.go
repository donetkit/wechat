package cache

import (
	"fmt"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	opts := &RedisOpts{
		Host:     "127.0.0.1",
		Port:     6379,
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

	d1, err := redis.XAdd("123456", "", []string{"", "troixtres21"}) //  []string{"tres", "troix"}
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(d1)
	dataXRead := redis.XRead("123456", 0)
	fmt.Println(dataXRead)
}
