package qrcode

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	context2 "github.com/donetkit/wechat/qqminiprogram/context"
	"github.com/donetkit/wechat/util"
)

const (
	createWXAQRCodeURL   = "https://api.q.qq.com/cgi-bin/wxaapp/createwxaqrcode?access_token=%s"
	getWXACodeURL        = "https://api.q.qq.com/wxa/getwxacode?access_token=%s"
	getWXACodeUnlimitURL = "https://api.q.qq.com/wxa/getwxacodeunlimit?access_token=%s"
)

// QRCode struct
type QRCode struct {
	*context2.Context
}

// NewQRCode 实例
func NewQRCode(context *context2.Context) *QRCode {
	qrCode := new(QRCode)
	qrCode.Context = context
	return qrCode
}

// Color QRCode color
type Color struct {
	R string `json:"r"`
	G string `json:"g"`
	B string `json:"b"`
}

// QRCoder 小程序码参数
type QRCoder struct {
	// page 必须是已经发布的小程序存在的页面,根路径前不要填加 /,不能携带参数（参数请放在scene字段里），如果不填写这个字段，默认跳主页面
	Page string `json:"page,omitempty"`
	// path 扫码进入的小程序页面路径
	Path string `json:"path,omitempty"`
	// width 图片宽度
	Width int `json:"width,omitempty"`
	// scene 最大32个可见字符，只支持数字，大小写英文以及部分特殊字符：!#$&'()*+,/:;=?@-._~，其它字符请自行编码为合法字符（因不支持%，中文无法使用 urlencode 处理，请使用其他编码方式）
	Scene string `json:"scene,omitempty"`
	// autoColor 自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调
	AutoColor bool `json:"auto_color,omitempty"`
	// lineColor AutoColor 为 false 时生效，使用 rgb 设置颜色 例如 {"r":"xxx","g":"xxx","b":"xxx"},十进制表示
	LineColor Color `json:"line_color,omitempty"`
	// isHyaline 是否需要透明底色
	IsHyaline bool `json:"is_hyaline,omitempty"`
}

// fetchCode 请求并返回二维码二进制数据
func (qrCode *QRCode) fetchCode(ctx context.Context, urlStr string, body interface{}) (response []byte, err error) {
	var accessToken string
	accessToken, err = qrCode.GetAccessTokenContext(ctx)
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
