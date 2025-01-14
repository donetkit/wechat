package apis

import (
	"encoding/json"
	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqDeleteParentUser 删除家长请求
// 文档：https://developer.work.weixin.qq.com/document/path/100152#删除家长
type ReqDeleteParentUser struct {
	UserId string `json:"userid"`
}

var _ urlValuer = ReqDeleteParentUser{}

func (x ReqDeleteParentUser) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespDeleteParentUser 删除家长响应
// 文档：https://developer.work.weixin.qq.com/document/path/100152#删除家长
type RespDeleteParentUser struct {
	CommonResp
}

var _ bodyer = RespDeleteParentUser{}

func (x RespDeleteParentUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecDeleteParentUser 删除家长
// 文档：https://developer.work.weixin.qq.com/document/path/100152#删除家长
func (c *ApiClient) ExecDeleteParentUser(req ReqDeleteParentUser) (RespDeleteParentUser, error) {
	var resp RespDeleteParentUser
	err := c.executeWXApiGet("/cgi-bin/school/user/delete_parent", req, &resp, true)
	if err != nil {
		return RespDeleteParentUser{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespDeleteParentUser{}, bizErr
	}
	return resp, nil
}
