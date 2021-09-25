package qqminiprogram

import (
	"github.com/donetkit/wechat/credential"
	"github.com/donetkit/wechat/qqminiprogram/analysis"
	"github.com/donetkit/wechat/qqminiprogram/auth"
	"github.com/donetkit/wechat/qqminiprogram/config"
	"github.com/donetkit/wechat/qqminiprogram/content"
	"github.com/donetkit/wechat/qqminiprogram/context"
	"github.com/donetkit/wechat/qqminiprogram/encryptor"
	"github.com/donetkit/wechat/qqminiprogram/message"
	"github.com/donetkit/wechat/qqminiprogram/qrcode"
	"github.com/donetkit/wechat/qqminiprogram/shortlink"
	"github.com/donetkit/wechat/qqminiprogram/subscribe"
	"github.com/donetkit/wechat/qqminiprogram/tcb"
	"github.com/donetkit/wechat/qqminiprogram/urllink"
	"github.com/donetkit/wechat/qqminiprogram/werun"
)

//QQMiniProgram 微信小程序相关API
type QQMiniProgram struct {
	ctx *context.Context
}

//NewQQMiniProgram 实例化小程序API
func NewQQMiniProgram(cfg *config.Config) *QQMiniProgram {
	defaultAkHandle := credential.NewDefaultAccessToken(cfg.AppID, cfg.AppSecret, credential.CacheKeyMiniProgramPrefix, cfg.Cache)
	ctx := &context.Context{
		Config:            cfg,
		AccessTokenHandle: defaultAkHandle,
	}
	return &QQMiniProgram{ctx}
}

//SetAccessTokenHandle 自定义access_token获取方式
func (miniProgram *QQMiniProgram) SetAccessTokenHandle(accessTokenHandle credential.AccessTokenHandle) {
	miniProgram.ctx.AccessTokenHandle = accessTokenHandle
}

// GetContext get Context
func (miniProgram *QQMiniProgram) GetContext() *context.Context {
	return miniProgram.ctx
}

// GetEncryptor  小程序加解密
func (miniProgram *QQMiniProgram) GetEncryptor() *encryptor.Encryptor {
	return encryptor.NewEncryptor(miniProgram.ctx)
}

//GetAuth 登录/用户信息相关接口
func (miniProgram *QQMiniProgram) GetAuth() *auth.Auth {
	return auth.NewAuth(miniProgram.ctx)
}

//GetAnalysis 数据分析
func (miniProgram *QQMiniProgram) GetAnalysis() *analysis.Analysis {
	return analysis.NewAnalysis(miniProgram.ctx)
}

//GetQRCode 小程序码相关API
func (miniProgram *QQMiniProgram) GetQRCode() *qrcode.QRCode {
	return qrcode.NewQRCode(miniProgram.ctx)
}

//GetTcb 小程序云开发API
func (miniProgram *QQMiniProgram) GetTcb() *tcb.Tcb {
	return tcb.NewTcb(miniProgram.ctx)
}

//GetSubscribe 小程序订阅消息
func (miniProgram *QQMiniProgram) GetSubscribe() *subscribe.Subscribe {
	return subscribe.NewSubscribe(miniProgram.ctx)
}

// GetCustomerMessage 客服消息接口
func (miniProgram *QQMiniProgram) GetCustomerMessage() *message.Manager {
	return message.NewCustomerMessageManager(miniProgram.ctx)
}

// GetWeRun 微信运动接口
func (miniProgram *QQMiniProgram) GetWeRun() *werun.WeRun {
	return werun.NewWeRun(miniProgram.ctx)
}

// GetContentSecurity 内容安全接口
func (miniProgram *QQMiniProgram) GetContentSecurity() *content.Content {
	return content.NewContent(miniProgram.ctx)
}

// GetURLLink 小程序URL Link接口
func (miniProgram *QQMiniProgram) GetURLLink() *urllink.URLLink {
	return urllink.NewURLLink(miniProgram.ctx)
}

// GetShortLink 小程序短链接口
func (miniProgram *QQMiniProgram) GetShortLink() *shortlink.ShortLink {
	return shortlink.NewShortLink(miniProgram.ctx)
}
