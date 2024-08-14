package qrcode

import (
	"context"
	"encoding/json"
	"fmt"
	openContext "github.com/donetkit/wechat/openplatform/context"
	"github.com/donetkit/wechat/util"
	"strings"
)

const (
	createWXAQRCodeURL   = "https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode?access_token=%s"
	getWXACodeURL        = "https://api.weixin.qq.com/wxa/getwxacode?access_token=%s"
	getWXACodeUnlimitURL = "https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=%s"
)

// QRCode struct
type QRCode struct {
	*openContext.Context
	appID string // 小程序appId
}

// NewQRCode 实例
func NewQRCode(context *openContext.Context, appID string) *QRCode {
	qrCode := new(QRCode)
	qrCode.Context = context
	qrCode.appID = appID
	return qrCode
}

// fetchCode 请求并返回二维码二进制数据
func (qrCode *QRCode) fetchCode(ctx context.Context, urlStr string, body interface{}) (response []byte, err error) {
	var accessToken string
	accessToken, err = qrCode.GetAuthAccessToken1(ctx, qrCode.appID)
	if err != nil {
		return
	}

	urlStr = fmt.Sprintf(urlStr, accessToken)
	var contentType string
	response, contentType, err = util.PostJSONWithRespContentType(urlStr, body)
	if err != nil {
		return
	}
	if strings.HasPrefix(contentType, "application/json") {
		// 返回错误信息
		var result util.CommonError
		err = json.Unmarshal(response, &result)
		if err == nil && result.ErrCode != 0 {
			err = fmt.Errorf("fetchCode error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
			return nil, err
		}
	}
	if contentType == "image/jpeg" {
		// 返回文件
		return response, nil
	}
	err = fmt.Errorf("fetchCode error : unknown response content type - %v", contentType)
	return nil, err
}

// CreateWXAQRCode 获取小程序二维码，适用于需要的码数量较少的业务场景
// 文档地址： https://developers.weixin.qq.com/miniprogram/dev/api/createWXAQRCode.html
func (qrCode *QRCode) CreateWXAQRCode(ctx context.Context, coderParams QRCoder) (response []byte, err error) {
	return qrCode.fetchCode(ctx, createWXAQRCodeURL, coderParams)
}

// GetWXACode 获取小程序码，适用于需要的码数量较少的业务场景
// 文档地址： https://developers.weixin.qq.com/miniprogram/dev/api/getWXACode.html
func (qrCode *QRCode) GetWXACode(ctx context.Context, coderParams QRCoder) (response []byte, err error) {
	return qrCode.fetchCode(ctx, getWXACodeURL, coderParams)
}

// GetWXACodeUnlimit 获取小程序码，适用于需要的码数量极多的业务场景
// 文档地址： https://developers.weixin.qq.com/miniprogram/dev/api/getWXACodeUnlimit.html
func (qrCode *QRCode) GetWXACodeUnlimit(ctx context.Context, coderParams QRCoder) (response []byte, err error) {
	return qrCode.fetchCode(ctx, getWXACodeUnlimitURL, coderParams)
}
