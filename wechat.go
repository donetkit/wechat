package wechat

import (
	"github.com/donetkit/contrib_cache/cache"
	"github.com/donetkit/wechat/qqminiprogram"
	qqMiniConfig "github.com/donetkit/wechat/qqminiprogram/config"

	"github.com/donetkit/wechat/miniprogram"
	miniConfig "github.com/donetkit/wechat/miniprogram/config"
	"github.com/donetkit/wechat/officialaccount"
	offConfig "github.com/donetkit/wechat/officialaccount/config"
	"github.com/donetkit/wechat/openplatform"
	openConfig "github.com/donetkit/wechat/openplatform/config"
	"github.com/donetkit/wechat/pay"
	payConfig "github.com/donetkit/wechat/pay/config"
)

// WeChat struct
type WeChat struct {
	cache cache.ICache
}

// NewWechat init
func NewWechat() *WeChat {
	return &WeChat{}
}

// SetCache 设置cache
func (wc *WeChat) SetCache(cache cache.ICache) {
	wc.cache = cache
}

// GetOfficialAccount 获取微信公众号实例
func (wc *WeChat) GetOfficialAccount(cfg *offConfig.Config) *officialaccount.OfficialAccount {
	if cfg.Cache == nil {
		cfg.Cache = wc.cache
	}
	return officialaccount.NewOfficialAccount(cfg)
}

// GetMiniProgram 获取小程序的实例
func (wc *WeChat) GetMiniProgram(cfg *miniConfig.Config) *miniprogram.MiniProgram {
	if cfg.Cache == nil {
		cfg.Cache = wc.cache
	}
	return miniprogram.NewMiniProgram(cfg)
}

// GetQQMiniProgram 获取小程序的实例
func (wc *WeChat) GetQQMiniProgram(cfg *qqMiniConfig.Config) *qqminiprogram.QQMiniProgram {
	if cfg.Cache == nil {
		cfg.Cache = wc.cache
	}
	return qqminiprogram.NewQQMiniProgram(cfg)
}

// GetPay 获取微信支付的实例
func (wc *WeChat) GetPay(cfg *payConfig.Config) *pay.Pay {
	return pay.NewPay(cfg)
}

// GetOpenPlatform 获取微信开放平台的实例
func (wc *WeChat) GetOpenPlatform(cfg *openConfig.Config) *openplatform.OpenPlatform {
	if cfg.Cache == nil {
		cfg.Cache = wc.cache
	}
	return openplatform.NewOpenPlatform(cfg)
}
