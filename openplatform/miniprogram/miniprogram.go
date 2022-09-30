package miniprogram

import (
	"context"
	"fmt"
	miniContext "github.com/donetkit/wechat/miniprogram/context"
	"github.com/donetkit/wechat/miniprogram/urllink"
	openContext "github.com/donetkit/wechat/openplatform/context"
	"github.com/donetkit/wechat/openplatform/miniprogram/auth"
	"github.com/donetkit/wechat/openplatform/miniprogram/basic"
	"github.com/donetkit/wechat/openplatform/miniprogram/component"
)

// MiniProgram 代小程序实现业务
type MiniProgram struct {
	AppID                  string
	openContext            *openContext.Context
	authorizerRefreshToken string
}

// NewMiniProgram 实例化
func NewMiniProgram(opCtx *openContext.Context, appID string) *MiniProgram {
	return &MiniProgram{
		openContext: opCtx,
		AppID:       appID,
	}
}

// GetAccessToken 获取ak
func (miniProgram *MiniProgram) GetAccessToken(ctx context.Context) (string, error) {
	ak, akErr := miniProgram.openContext.GetAuthrAccessToken(ctx, miniProgram.AppID)
	if akErr == nil {
		return ak, nil
	}
	if miniProgram.authorizerRefreshToken == "" {
		return "", fmt.Errorf("please set the authorizer_refresh_token first")
	}
	akRes, akResErr := miniProgram.GetComponent().RefreshAuthrToken(ctx, miniProgram.AppID, miniProgram.authorizerRefreshToken)
	if akResErr != nil {
		return "", akResErr
	}
	return akRes.AccessToken, nil
}

// SetAuthorizerRefreshToken 设置代执操作业务授权账号authorizer_refresh_token
func (miniProgram *MiniProgram) SetAuthorizerRefreshToken(authorizerRefreshToken string) *MiniProgram {
	miniProgram.authorizerRefreshToken = authorizerRefreshToken
	return miniProgram
}

// GetComponent get component
// 快速注册小程序相关
func (miniProgram *MiniProgram) GetComponent() *component.Component {
	return component.NewComponent(miniProgram.openContext)
}

// GetBasic 基础信息设置
func (miniProgram *MiniProgram) GetBasic() *basic.Basic {
	return basic.NewBasic(miniProgram.openContext, miniProgram.AppID)
}

// GetBasic 基础信息设置
func (miniProgram *MiniProgram) GetAuth() *auth.Auth {
	return auth.NewAuth(miniProgram.openContext, miniProgram.AppID)
}

// GetURLLink 小程序URL Link接口 调用前需确认已调用 SetAuthorizerRefreshToken 避免由于缓存中 authorizer_access_token 过期执行中断
func (miniProgram *MiniProgram) GetURLLink() *urllink.URLLink {
	return urllink.NewURLLink(&miniContext.Context{
		AccessTokenHandle: miniProgram,
	})
}
