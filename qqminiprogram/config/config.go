// Package config 小程序config配置
package config

import (
	"github.com/donetkit/contrib_cache/cache"
)

// Config .config for 小程序
type Config struct {
	AppID       string `json:"app_id"`     // appid
	AppSecret   string `json:"app_secret"` // appsecret
	Cache       cache.ICache
	UseStableAK bool // use the stable access_token
}
