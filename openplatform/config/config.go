package config

import (
	"github.com/donetkit/wechat/cache"
	"sync"
)

//Config config for 微信开放平台
type Config struct {
	AppID          string `json:"app_id"`           //appid 开放平台
	AppSecret      string `json:"app_secret"`       //appsecret 开放平台
	Token          string `json:"token"`            //token 开放平台
	EncodingAESKey string `json:"encoding_aes_key"` //EncodingAESKey 开放平台
	Cache          cache.Cache
	Lock           *sync.Mutex
}
