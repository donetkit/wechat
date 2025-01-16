package js

import (
	"context"
	"fmt"

	"github.com/donetkit/wechat/credential"
	context2 "github.com/donetkit/wechat/officialaccount/context"
	officialJs "github.com/donetkit/wechat/officialaccount/js"
	"github.com/donetkit/wechat/util"
)

// Js wx jssdk
type Js struct {
	*context2.Context
	credential.JsTicketHandle
}

// NewJs init
func NewJs(context *context2.Context, appID string) *Js {
	js := new(Js)
	js.Context = context
	jsTicketHandle := credential.NewDefaultJsTicket(appID, credential.CacheKeyOfficialAccountPrefix, context.Cache)
	js.SetJsTicketHandle(jsTicketHandle)
	return js
}

// SetJsTicketHandle 自定义js ticket取值方式
func (js *Js) SetJsTicketHandle(ticketHandle credential.JsTicketHandle) {
	js.JsTicketHandle = ticketHandle
}

// GetConfig 第三方平台 - 获取jssdk需要的配置参数
// uri 为当前网页地址
func (js *Js) GetConfig(ctx context.Context, uri, appid string) (config *officialJs.Config, err error) {
	config = new(officialJs.Config)
	var accessToken string
	accessToken, err = js.GetAccessTokenContext(ctx)
	if err != nil {
		return
	}
	var ticketStr string
	ticketStr, err = js.GetTicket(ctx, accessToken)
	if err != nil {
		return
	}

	nonceStr := util.RandomStr(16)
	timestamp := util.GetCurrTS()
	str := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s", ticketStr, nonceStr, timestamp, uri)
	sigStr := util.Signature(str)

	config.AppID = appid
	config.NonceStr = nonceStr
	config.Timestamp = timestamp
	config.Signature = sigStr
	return
}
