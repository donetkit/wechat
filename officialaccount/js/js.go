package js

import (
	"context"
	"fmt"

	"github.com/donetkit/wechat/credential"
	context2 "github.com/donetkit/wechat/officialaccount/context"
	"github.com/donetkit/wechat/util"
)

// Js struct
type Js struct {
	*context2.Context
	credential.JsTicketHandle
}

// Config 返回给用户jssdk配置信息
type Config struct {
	AppID     string `json:"app_id"`
	Timestamp int64  `json:"timestamp"`
	NonceStr  string `json:"nonce_str"`
	Signature string `json:"signature"`
}

//NewJs init
func NewJs(context *context2.Context) *Js {
	js := new(Js)
	js.Context = context
	jsTicketHandle := credential.NewDefaultJsTicket(context.AppID, credential.CacheKeyOfficialAccountPrefix, context.Cache)
	js.SetJsTicketHandle(jsTicketHandle)
	return js
}

//SetJsTicketHandle 自定义js ticket取值方式
func (js *Js) SetJsTicketHandle(ticketHandle credential.JsTicketHandle) {
	js.JsTicketHandle = ticketHandle
}

//GetConfig 获取jssdk需要的配置参数
//uri 为当前网页地址
func (js *Js) GetConfig(ctx context.Context, uri string) (config *Config, err error) {
	config = new(Config)
	var accessToken string
	accessToken, err = js.GetAccessToken(ctx)
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

	config.AppID = js.AppID
	config.NonceStr = nonceStr
	config.Timestamp = timestamp
	config.Signature = sigStr
	return
}
