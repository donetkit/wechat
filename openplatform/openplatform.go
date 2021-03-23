package openplatform

import (
	"github.com/donetkit/wechat/officialaccount/server"
	"github.com/donetkit/wechat/openplatform/account"
	"github.com/donetkit/wechat/openplatform/config"
	"github.com/donetkit/wechat/openplatform/context"
	"github.com/donetkit/wechat/openplatform/miniprogram"
	"github.com/donetkit/wechat/openplatform/officialaccount"
	"github.com/gin-gonic/gin"
)

//OpenPlatform 微信开放平台相关api
type OpenPlatform struct {
	*context.Context
}

//NewOpenPlatform new openplatform
func NewOpenPlatform(cfg *config.Config) *OpenPlatform {
	if cfg.Cache == nil {
		panic("cache 未设置")
	}
	ctx := &context.Context{
		Config: cfg,
	}
	return &OpenPlatform{ctx}
}

//GetServer get server
func (openPlatform *OpenPlatform) GetServer(c *gin.Context, appID string) *server.Server {
	off := officialaccount.NewOfficialAccount(openPlatform.Context, appID)
	return off.GetServer(c)
}

//GetOfficialAccount 公众号代处理
func (openPlatform *OpenPlatform) GetOfficialAccount(appID string) *officialaccount.OfficialAccount {
	return officialaccount.NewOfficialAccount(openPlatform.Context, appID)
}

//GetMiniProgram 小程序代理
func (openPlatform *OpenPlatform) GetMiniProgram(appID string) *miniprogram.MiniProgram {
	return miniprogram.NewMiniProgram(openPlatform.Context, appID)
}

//GetAccountManager 账号管理
//TODO
func (openPlatform *OpenPlatform) GetAccountManager() *account.Account {
	return account.NewAccount(openPlatform.Context)
}
