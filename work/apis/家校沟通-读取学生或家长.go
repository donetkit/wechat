package apis

import (
	"encoding/json"
	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetUserSchool 读取学生或家长请求
// 文档：https://developer.work.weixin.qq.com/document/path/96738#读取学生或家长
type ReqGetUserSchool struct {
	UserId string `json:"userid"`
}

var _ urlValuer = ReqGetUserSchool{}

func (x ReqGetUserSchool) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespGetUserSchool 读取学生或家长响应
// 文档：https://developer.work.weixin.qq.com/document/path/96738#读取学生或家长
type RespGetUserSchool struct {
	CommonResp
	UserType int `json:"user_type"`
	Student  struct {
		StudentUserid string `json:"student_userid"`
		Name          string `json:"name"`
		Department    []int  `json:"department"`
		Parents       []struct {
			ParentUserid   string `json:"parent_userid"`
			Relation       string `json:"relation"`
			Mobile         string `json:"mobile"`
			IsSubscribe    int    `json:"is_subscribe"`
			ExternalUserid string `json:"external_userid,omitempty"`
		} `json:"parents"`
	} `json:"student"`
	Parent struct {
		ParentUserid   string `json:"parent_userid"`
		Mobile         string `json:"mobile"`
		IsSubscribe    int    `json:"is_subscribe"`
		ExternalUserid string `json:"external_userid"`
		Children       []struct {
			StudentUserid string `json:"student_userid"`
			Relation      string `json:"relation"`
		} `json:"children"`
	} `json:"parent"`
}

var _ bodyer = RespGetUserSchool{}

func (x RespGetUserSchool) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetUserSchool 读取学生或家长
// 文档：https://developer.work.weixin.qq.com/document/path/96738#读取学生或家长
func (c *ApiClient) ExecGetUserSchool(req ReqGetUserSchool) (RespGetUserSchool, error) {
	var resp RespGetUserSchool
	err := c.executeWXApiGet("/cgi-bin/school/user/get", req, &resp, true)
	if err != nil {
		return RespGetUserSchool{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetUserSchool{}, bizErr
	}
	return resp, nil
}
