package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetuserdetailCustomizeService 获取访问用户敏感信息请求
// 文档：https://developer.work.weixin.qq.com/document/path/91122#获取访问用户敏感信息
type ReqGetuserdetailCustomizeService struct {
	// UserTicket 成员票据，必填
	UserTicket string `json:"user_ticket"`
}

var _ bodyer = ReqGetuserdetailCustomizeService{}

func (x ReqGetuserdetailCustomizeService) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetuserdetailCustomizeService 获取访问用户敏感信息响应
// 文档：https://developer.work.weixin.qq.com/document/path/96443#获取访问用户敏感信息
type RespGetuserdetailCustomizeService struct {
	Avatar string `json:"avatar"`
	Corpid string `json:"corpid"`
	CommonResp
	Gender  string `json:"gender"`
	Name    string `json:"name"`
	QrCode  string `json:"qr_code"`
	Userid  string `json:"userid"`
	Mobile  string `json:"mobile"`
	Email   string `json:"email"`
	BizMail string `json:"biz_mail"`
	Address string `json:"address"`
}

var _ bodyer = RespGetuserdetailCustomizeService{}

func (x RespGetuserdetailCustomizeService) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetuserdetailCustomizeService 获取访问用户敏感信息
// 文档：https://developer.work.weixin.qq.com/document/path/96443#获取访问用户敏感信息
func (c *ApiClient) ExecGetuserdetailCustomizeService(req ReqGetuserdetailCustomizeService) (RespGetuserdetailCustomizeService, error) {
	var resp RespGetuserdetailCustomizeService
	err := c.executeWXApiPost("/cgi-bin/auth/getuserdetail", req, &resp, true)
	if err != nil {
		return RespGetuserdetailCustomizeService{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetuserdetailCustomizeService{}, bizErr
	}
	return resp, nil
}
