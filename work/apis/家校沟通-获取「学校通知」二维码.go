package apis

import (
	"encoding/json"
	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetSubscribeQrCodeExternalcontact 获取「学校通知」二维码请求
// 文档：https://developer.work.weixin.qq.com/document/path/96719#获取「学校通知」二维码
type ReqGetSubscribeQrCodeExternalcontact struct{}

var _ urlValuer = ReqGetSubscribeQrCodeExternalcontact{}

func (x ReqGetSubscribeQrCodeExternalcontact) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespGetSubscribeQrCodeExternalcontact 获取「学校通知」二维码响应
// 文档：https://developer.work.weixin.qq.com/document/path/96719#获取「学校通知」二维码
type RespGetSubscribeQrCodeExternalcontact struct {
	CommonResp
	QrcodeBig    string `json:"qrcode_big"`
	QrcodeMiddle string `json:"qrcode_middle"`
	QrcodeThumb  string `json:"qrcode_thumb"`
}

var _ bodyer = RespGetSubscribeQrCodeExternalcontact{}

func (x RespGetSubscribeQrCodeExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetSubscribeQrCodeExternalcontact 获取「学校通知」二维码
// 文档：https://developer.work.weixin.qq.com/document/path/96719#获取「学校通知」二维码
func (c *ApiClient) ExecGetSubscribeQrCodeExternalcontact(req ReqGetSubscribeQrCodeExternalcontact) (RespGetSubscribeQrCodeExternalcontact, error) {
	var resp RespGetSubscribeQrCodeExternalcontact
	err := c.executeWXApiGet("/cgi-bin/externalcontact/get_subscribe_qr_code", req, &resp, true)
	if err != nil {
		return RespGetSubscribeQrCodeExternalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetSubscribeQrCodeExternalcontact{}, bizErr
	}
	return resp, nil
}
