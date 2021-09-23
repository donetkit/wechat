package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

//Redis redis cache
type Redis struct {
	client *redis.Client
}

//RedisOpts redis 连接属性
type RedisOpts struct {
	Host        string `yml:"host" json:"host"`
	Port        int    `yml:"port" json:"port"`
	Password    string `yml:"password" json:"password"`
	Database    int    `yml:"database" json:"database"`
	MaxIdle     int    `yml:"max_idle" json:"max_idle"`
	MaxActive   int    `yml:"max_active" json:"max_active"`
	IdleTimeout int    `yml:"idle_timeout" json:"idle_timeout"` // second
}

var ctx = context.Background()

//NewRedis 实例化
func NewRedis(opts *RedisOpts) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", opts.Host, opts.Port),
		Password: opts.Password, // no password set
		DB:       opts.Database, // use default DB
	})
	return &Redis{client}
}

//Get 获取一个值
func (r *Redis) Get(key string) interface{} {
	data, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		panic(err)
	}
	var reply interface{}
	if err = json.Unmarshal(data, &reply); err != nil {
		return nil
	}
	return reply
}

//Set 设置一个值
func (r *Redis) Set(key string, val interface{}, timeout time.Duration) (err error) {
	var data []byte
	if data, err = json.Marshal(val); err != nil {
		return
	}
	err = r.client.Set(ctx, key, data, timeout).Err()
	if err != nil {
		panic(err)
	}
	return
}

//IsExist 判断key是否存在
func (r *Redis) IsExist(key string) bool {
	i := r.client.Exists(ctx, key).Val()
	return i > 0
}

//Delete 删除
func (r *Redis) Delete(key string) error {
	err := r.client.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}

//XRead
func (r *Redis) XRead(key string, count int64) []redis.XStream {
	return r.client.XReadStreams(ctx, key, fmt.Sprintf("%d", count)).Val()
}

// XAdd
func (r *Redis) XAdd(key, id string, values []string) (string, error) {
	id, err := r.client.XAdd(ctx, &redis.XAddArgs{
		Stream: key,
		ID:     id,
		Values: values,
	}).Result()
	if err != nil {
		return "", err
	}
	return id, nil
}

// XDel
func (r *Redis) XDel(key string, id string) (int64, error) {
	n, err := r.client.XDel(ctx, key, id).Result()
	if err != nil {
		return 0, err
	}
	return n, nil
}
