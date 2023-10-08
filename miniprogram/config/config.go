// Package config 小程序 config 配置
package config

import "github.com/donetkit/contrib/utils/cache"

// Config .config for 小程序
type Config struct {
	AppID     string `json:"app_id"`     // appid
	AppSecret string `json:"app_secret"` // appSecret
	AppKey    string `json:"app_key"`    // appKey
	Cache     cache.ICache
}
