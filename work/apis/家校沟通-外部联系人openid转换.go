package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqConvertToOpenidExternalcontact 外部联系人openid转换请求
// 文档：https://developer.work.weixin.qq.com/document/path/96721#外部联系人openid转换
type ReqConvertToOpenidExternalcontact struct {
	ExternalUserid string `json:"external_userid"` // 外部联系人的userid，注意不是企业成员的账号
}

var _ bodyer = ReqConvertToOpenidExternalcontact{}

func (x ReqConvertToOpenidExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespConvertToOpenidExternalcontact 外部联系人openid转换响应
// 文档：https://developer.work.weixin.qq.com/document/path/96721#外部联系人openid转换
type RespConvertToOpenidExternalcontact struct {
	CommonResp
	Openid string `json:"openid"` // 该企业的外部联系人openid
}

var _ bodyer = RespConvertToOpenidExternalcontact{}

func (x RespConvertToOpenidExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecConvertToOpenidExternalcontact 外部联系人openid转换 用于调用支付相关接口
// 文档：https://developer.work.weixin.qq.com/document/path/96721#外部联系人openid转换
func (c *ApiClient) ExecConvertToOpenidExternalcontact(req ReqConvertToOpenidExternalcontact) (RespConvertToOpenidExternalcontact, error) {
	var resp RespConvertToOpenidExternalcontact
	err := c.executeWXApiPost(":/cgi-bin/externalcontact/convert_to_openid", req, &resp, true)
	if err != nil {
		return RespConvertToOpenidExternalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespConvertToOpenidExternalcontact{}, bizErr
	}
	return resp, nil
}
