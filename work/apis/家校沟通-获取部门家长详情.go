package apis

import (
	"encoding/json"
	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqListParentUser 获取部门家长详情请求
// 文档：https://developer.work.weixin.qq.com/document/path/96741#获取部门家长详情
type ReqListParentUser struct {
	DepartmentID string `json:"department_id"`
}

var _ urlValuer = ReqListParentUser{}

func (x ReqListParentUser) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespListParentUser 获取部门家长详情响应
// 文档：https://developer.work.weixin.qq.com/document/path/96741#获取部门家长详情
type RespListParentUser struct {
	CommonResp
	Parents []struct {
		ParentUserid   string `json:"parent_userid"`
		Mobile         string `json:"mobile"`
		IsSubscribe    int    `json:"is_subscribe"`
		ExternalUserid string `json:"external_userid,omitempty"`
		Children       []struct {
			StudentUserid string `json:"student_userid"`
			Relation      string `json:"relation"`
			Name          string `json:"name"`
		} `json:"children"`
	} `json:"parents"`
}

var _ bodyer = RespListParentUser{}

func (x RespListParentUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecListParentUser 获取部门家长详情
// 文档：https://developer.work.weixin.qq.com/document/path/96741#获取部门家长详情
func (c *ApiClient) ExecListParentUser(req ReqListParentUser) (RespListParentUser, error) {
	var resp RespListParentUser
	err := c.executeWXApiGet("/cgi-bin/school/user/list_parent", req, &resp, true)
	if err != nil {
		return RespListParentUser{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespListParentUser{}, bizErr
	}
	return resp, nil
}
