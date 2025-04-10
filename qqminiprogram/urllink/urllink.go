package urllink

import (
	"context"
	"fmt"

	context2 "github.com/donetkit/wechat/qqminiprogram/context"
	"github.com/donetkit/wechat/util"
)

// URLLink 小程序 URL Link
type URLLink struct {
	*context2.Context
}

// NewURLLink 实例化
func NewURLLink(ctx *context2.Context) *URLLink {
	return &URLLink{Context: ctx}
}

const generateURL = "https://api.q.qq.com/wxa/generate_urllink"

// TExpireType 失效类型 (指定时间戳/指定间隔)
type TExpireType int

const (
	// ExpireTypeTime 指定时间戳后失效
	ExpireTypeTime TExpireType = 0

	// ExpireTypeInterval 间隔指定天数后失效
	ExpireTypeInterval TExpireType = 1
)

// ULParams 请求参数
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-link/urllink.generate.html#请求参数
type ULParams struct {
	Path           string      `json:"path"`
	Query          string      `json:"query"`
	IsExpire       bool        `json:"is_expire"`
	ExpireType     TExpireType `json:"expire_type"`
	ExpireTime     int64       `json:"expire_time"`
	ExpireInterval int         `json:"expire_interval"`
}

// ULResult 返回的结果
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-link/urllink.generate.html#返回值
type ULResult struct {
	util.CommonError

	URLLink string `json:"url_link"`
}

// Generate 生成url link
func (u *URLLink) Generate(ctx context.Context, params *ULParams) (string, error) {
	var accessToken string
	accessToken, err := u.GetAccessTokenContext(ctx)
	if err != nil {
		return "", err
	}

	uri := fmt.Sprintf("%s?access_token=%s", generateURL, accessToken)
	response, err := util.PostJSONContext(ctx, uri, params)
	if err != nil {
		return "", err
	}
	var resp ULResult
	err = util.DecodeWithError(response, &resp, "URLLink.Generate")
	return resp.URLLink, err
}
