package config

import (
	"github.com/donetkit/contrib_cache/cache"
)

// Config config for 微信公众号
type Config struct {
	AppID          string `json:"app_id"`           //appid
	AppSecret      string `json:"app_secret"`       //appsecret
	Token          string `json:"token"`            //token
	EncodingAESKey string `json:"encoding_aes_key"` //EncodingAESKey
	Cache          cache.ICache
	UseStableAK    bool // use the stable access_token
}
