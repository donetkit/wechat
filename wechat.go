package wechat

import (
	"github.com/donetkit/wechat/qqminiprogram"
	qqMiniConfig "github.com/donetkit/wechat/qqminiprogram/config"

	"github.com/donetkit/contrib/utils/cache"
	"github.com/donetkit/wechat/miniprogram"
	miniConfig "github.com/donetkit/wechat/miniprogram/config"
	"github.com/donetkit/wechat/officialaccount"
	offConfig "github.com/donetkit/wechat/officialaccount/config"
	"github.com/donetkit/wechat/openplatform"
	openConfig "github.com/donetkit/wechat/openplatform/config"
	"github.com/donetkit/wechat/pay"
	payConfig "github.com/donetkit/wechat/pay/config"
	//workConfig "github.com/donetkit/wechat/work/config"
)

// Wechat struct
type Wechat struct {
	cache cache.ICache
}

// NewWechat init
func NewWechat() *Wechat {
	return &Wechat{}
}

//SetCache 设置cache
func (wc *Wechat) SetCache(cahce cache.ICache) {
	wc.cache = cahce
}

//GetOfficialAccount 获取微信公众号实例
func (wc *Wechat) GetOfficialAccount(cfg *offConfig.Config) *officialaccount.OfficialAccount {
	if cfg.Cache == nil {
		cfg.Cache = wc.cache
	}
	return officialaccount.NewOfficialAccount(cfg)
}

// GetMiniProgram 获取小程序的实例
func (wc *Wechat) GetMiniProgram(cfg *miniConfig.Config) *miniprogram.MiniProgram {
	if cfg.Cache == nil {
		cfg.Cache = wc.cache
	}
	return miniprogram.NewMiniProgram(cfg)
}

// GetQQMiniProgram 获取小程序的实例
func (wc *Wechat) GetQQMiniProgram(cfg *qqMiniConfig.Config) *qqminiprogram.QQMiniProgram {
	if cfg.Cache == nil {
		cfg.Cache = wc.cache
	}
	return qqminiprogram.NewQQMiniProgram(cfg)
}

// GetPay 获取微信支付的实例
func (wc *Wechat) GetPay(cfg *payConfig.Config) *pay.Pay {
	return pay.NewPay(cfg)
}

// GetOpenPlatform 获取微信开放平台的实例
func (wc *Wechat) GetOpenPlatform(cfg *openConfig.Config) *openplatform.OpenPlatform {
	return openplatform.NewOpenPlatform(cfg)
}

// GetWork 获取企业微信的实例
//func (wc *Wechat) GetWork(cfg *workConfig.Config) *work.Work {
//	return work.NewWork(cfg)
//}
