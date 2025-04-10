package miniprogram

import (
	"github.com/donetkit/wechat/credential"
	"github.com/donetkit/wechat/internal/openapi"
	"github.com/donetkit/wechat/miniprogram/analysis"
	"github.com/donetkit/wechat/miniprogram/auth"
	"github.com/donetkit/wechat/miniprogram/business"
	"github.com/donetkit/wechat/miniprogram/config"
	"github.com/donetkit/wechat/miniprogram/content"
	"github.com/donetkit/wechat/miniprogram/context"
	"github.com/donetkit/wechat/miniprogram/encryptor"
	"github.com/donetkit/wechat/miniprogram/message"
	"github.com/donetkit/wechat/miniprogram/minidrama"
	"github.com/donetkit/wechat/miniprogram/order"
	"github.com/donetkit/wechat/miniprogram/privacy"
	"github.com/donetkit/wechat/miniprogram/qrcode"
	"github.com/donetkit/wechat/miniprogram/redpacketcover"
	"github.com/donetkit/wechat/miniprogram/riskcontrol"
	"github.com/donetkit/wechat/miniprogram/security"
	"github.com/donetkit/wechat/miniprogram/shortlink"
	"github.com/donetkit/wechat/miniprogram/subscribe"
	"github.com/donetkit/wechat/miniprogram/tcb"
	"github.com/donetkit/wechat/miniprogram/urllink"
	"github.com/donetkit/wechat/miniprogram/virtualpayment"
	"github.com/donetkit/wechat/miniprogram/werun"
)

// MiniProgram 微信小程序相关API
type MiniProgram struct {
	ctx *context.Context
}

// NewMiniProgram 实例化小程序API
func NewMiniProgram(cfg *config.Config) *MiniProgram {
	//defaultAkHandle := credential.NewDefaultAccessToken(cfg.AppID, cfg.AppSecret, credential.CacheKeyMiniProgramPrefix, cfg.Cache)
	var defaultAkHandle credential.AccessTokenContextHandle
	const cacheKeyPrefix = credential.CacheKeyMiniProgramPrefix
	if cfg.UseStableAK {
		defaultAkHandle = credential.NewStableAccessToken(cfg.AppID, cfg.AppSecret, cacheKeyPrefix, cfg.Cache)
	} else {
		defaultAkHandle = credential.NewDefaultAccessToken(cfg.AppID, cfg.AppSecret, cacheKeyPrefix, cfg.Cache)
	}
	ctx := &context.Context{
		Config:                   cfg,
		AccessTokenContextHandle: defaultAkHandle,
	}
	return &MiniProgram{ctx}
}

// SetAccessTokenHandle 自定义access_token获取方式
func (miniProgram *MiniProgram) SetAccessTokenHandle(accessTokenHandle credential.AccessTokenContextHandle) {
	miniProgram.ctx.AccessTokenContextHandle = accessTokenHandle
}

// GetContext get Context
func (miniProgram *MiniProgram) GetContext() *context.Context {
	return miniProgram.ctx
}

// GetEncryptor  小程序加解密
func (miniProgram *MiniProgram) GetEncryptor() *encryptor.Encryptor {
	return encryptor.NewEncryptor(miniProgram.ctx)
}

// GetAuth 登录/用户信息相关接口
func (miniProgram *MiniProgram) GetAuth() *auth.Auth {
	return auth.NewAuth(miniProgram.ctx)
}

// GetAnalysis 数据分析
func (miniProgram *MiniProgram) GetAnalysis() *analysis.Analysis {
	return analysis.NewAnalysis(miniProgram.ctx)
}

// GetBusiness 业务接口
func (miniProgram *MiniProgram) GetBusiness() *business.Business {
	return business.NewBusiness(miniProgram.ctx)
}

// GetPrivacy 小程序隐私协议相关API
func (miniProgram *MiniProgram) GetPrivacy() *privacy.Privacy {
	return privacy.NewPrivacy(miniProgram.ctx)
}

// GetQRCode 小程序码相关API
func (miniProgram *MiniProgram) GetQRCode() *qrcode.QRCode {
	return qrcode.NewQRCode(miniProgram.ctx)
}

// GetTcb 小程序云开发API
func (miniProgram *MiniProgram) GetTcb() *tcb.Tcb {
	return tcb.NewTcb(miniProgram.ctx)
}

// GetSubscribe 小程序订阅消息
func (miniProgram *MiniProgram) GetSubscribe() *subscribe.Subscribe {
	return subscribe.NewSubscribe(miniProgram.ctx)
}

// GetCustomerMessage 客服消息接口
func (miniProgram *MiniProgram) GetCustomerMessage() *message.Manager {
	return message.NewCustomerMessageManager(miniProgram.ctx)
}

// GetWeRun 微信运动接口
func (miniProgram *MiniProgram) GetWeRun() *werun.WeRun {
	return werun.NewWeRun(miniProgram.ctx)
}

// GetContentSecurity 内容安全接口
func (miniProgram *MiniProgram) GetContentSecurity() *content.Content {
	return content.NewContent(miniProgram.ctx)
}

// GetURLLink 小程序URL Link接口
func (miniProgram *MiniProgram) GetURLLink() *urllink.URLLink {
	return urllink.NewURLLink(miniProgram.ctx)
}

// GetRiskControl 安全风控接口
func (miniProgram *MiniProgram) GetRiskControl() *riskcontrol.RiskControl {
	return riskcontrol.NewRiskControl(miniProgram.ctx)
}

// GetSecurity 内容安全接口
func (miniProgram *MiniProgram) GetSecurity() *security.Security {
	return security.NewSecurity(miniProgram.ctx)
}

// GetShortLink 小程序短链接口
func (miniProgram *MiniProgram) GetShortLink() *shortlink.ShortLink {
	return shortlink.NewShortLink(miniProgram.ctx)
}

// GetOpenAPI openApi管理接口
func (miniProgram *MiniProgram) GetOpenAPI() *openapi.OpenAPI {
	return openapi.NewOpenAPI(miniProgram.ctx)
}

// GetVirtualPayment 小程序虚拟支付
func (miniProgram *MiniProgram) GetVirtualPayment() *virtualpayment.VirtualPayment {
	return virtualpayment.NewVirtualPayment(miniProgram.ctx)
}

// GetMessageReceiver 获取消息推送接收器
func (miniProgram *MiniProgram) GetMessageReceiver() *message.PushReceiver {
	return message.NewPushReceiver(miniProgram.ctx)
}

// GetShipping 小程序发货信息管理服务
func (miniProgram *MiniProgram) GetShipping() *order.Shipping {
	return order.NewShipping(miniProgram.ctx)
}

// GetMiniDrama 小程序娱乐微短剧
func (miniProgram *MiniProgram) GetMiniDrama() *minidrama.MiniDrama {
	return minidrama.NewMiniDrama(miniProgram.ctx)
}

// GetRedPacketCover 小程序微信红包封面 API
func (miniProgram *MiniProgram) GetRedPacketCover() *redpacketcover.RedPacketCover {
	return redpacketcover.NewRedPacketCover(miniProgram.ctx)
}

// GetUpdatableMessage 小程序动态消息
func (miniProgram *MiniProgram) GetUpdatableMessage() *message.UpdatableMessage {
	return message.NewUpdatableMessage(miniProgram.ctx)
}
