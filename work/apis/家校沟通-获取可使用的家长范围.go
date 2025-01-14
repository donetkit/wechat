package apis

import (
	"encoding/json"
	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetAllowScopeAgent 获取可使用的家长范围请求
// 文档：https://developer.work.weixin.qq.com/document/path/96725#获取可使用的家长范围
type ReqGetAllowScopeAgent struct {
	AgentId string `json:"agentid"`
}

var _ urlValuer = ReqGetAllowScopeAgent{}

func (x ReqGetAllowScopeAgent) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespGetAllowScopeAgent 获取可使用的家长范围响应
// 文档：https://developer.work.weixin.qq.com/document/path/96725#获取可使用的家长范围
type RespGetAllowScopeAgent struct {
	CommonResp
	AllowScope struct {
		Students struct {
			Userid []string `json:"userid"` // 家长可在微信「学校通知-学校应用」使用该应用的学生列表
		} `json:"students"`
		Departments struct {
			Partyid []int `json:"partyid"` // 家长可在微信「学校通知-学校应用」使用该应用的部门列表
		} `json:"departments"`
	} `json:"allow_scope"`
}

var _ bodyer = RespGetAllowScopeAgent{}

func (x RespGetAllowScopeAgent) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetAllowScopeAgent 获取可使用的家长范围
// 文档：https://developer.work.weixin.qq.com/document/path/96725#获取可使用的家长范围
func (c *ApiClient) ExecGetAllowScopeAgent(req ReqGetAllowScopeAgent) (RespGetAllowScopeAgent, error) {
	var resp RespGetAllowScopeAgent
	err := c.executeWXApiGet("/cgi-bin/school/agent/get_allow_scope", req, &resp, true)
	if err != nil {
		return RespGetAllowScopeAgent{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetAllowScopeAgent{}, bizErr
	}
	return resp, nil
}
