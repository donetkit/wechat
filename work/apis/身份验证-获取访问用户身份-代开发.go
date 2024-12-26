package apis

import (
	"encoding/json"

	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetuserinfoCustomizeService 获取访问用户身份请求
// 文档：https://developer.work.weixin.qq.com/document/path/96442#获取访问用户身份
type ReqGetuserinfoCustomizeService struct {
	// Code 通过成员授权获取到的code，最大为512字节。每次成员授权带上的code将不一样，code只能使用一次，5分钟未被使用自动过期。，必填
	Code string `json:"code"`
}

var _ urlValuer = ReqGetuserinfoCustomizeService{}

func (x ReqGetuserinfoCustomizeService) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespGetuserinfoCustomizeService 获取访问用户身份响应
// 文档：https://developer.work.weixin.qq.com/document/path/96442#获取访问用户身份
type RespGetuserinfoCustomizeService struct {
	CommonResp
	CorpID     string `json:"CorpId"`
	UserID     string `json:"UserId"`
	DeviceID   string `json:"DeviceId"`
	UserTicket string `json:"user_ticket"`
	ExpiresIn  int    `json:"expires_in"`
	OpenUserID string `json:"open_userid"`
}

var _ bodyer = RespGetuserinfoCustomizeService{}

func (x RespGetuserinfoCustomizeService) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetuserinfoCustomizeService 获取访问用户身份
// 文档：https://developer.work.weixin.qq.com/document/path/91121#获取访问用户身份
func (c *ApiClient) ExecGetuserinfoCustomizeService(req ReqGetuserinfoCustomizeService) (RespGetuserinfoCustomizeService, error) {
	var resp RespGetuserinfoCustomizeService
	err := c.executeWXApiGet("/cgi-bin/auth/getuserinfo", req, &resp, true)
	if err != nil {
		return RespGetuserinfoCustomizeService{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetuserinfoCustomizeService{}, bizErr
	}
	return resp, nil
}
