package apis

import (
	"encoding/json"
	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqListDepartmentSchool 获取部门列表请求
// 文档：https://developer.work.weixin.qq.com/document/path/96745#获取部门列表
type ReqListDepartmentSchool struct {
	ID int `json:"id"`
}

var _ urlValuer = ReqListDepartmentSchool{}

func (x ReqListDepartmentSchool) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespListDepartmentSchool 获取部门列表响应
// 文档：https://developer.work.weixin.qq.com/document/path/96745#获取部门列表
var _ bodyer = RespListDepartmentSchool{}

func (x RespListDepartmentSchool) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecListDepartmentSchool 获取部门列表
// 文档：https://developer.work.weixin.qq.com/document/path/96745#获取部门列表
func (c *ApiClient) ExecListDepartmentSchool(req ReqListDepartmentSchool) (RespListDepartmentSchool, error) {
	var resp RespListDepartmentSchool
	err := c.executeWXApiGet("/cgi-bin/school/department/list", req, &resp, true)
	if err != nil {
		return RespListDepartmentSchool{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespListDepartmentSchool{}, bizErr
	}
	return resp, nil
}

type RespListDepartmentSchool struct {
	CommonResp
	Departments []struct {
		Name             string `json:"name"`
		Parentid         int    `json:"parentid"`
		Id               int    `json:"id"`
		DepartmentAdmins []struct {
			Userid  string `json:"userid"`
			Type    int    `json:"type"`
			Subject string `json:"subject,omitempty"`
		} `json:"department_admins"`
		Type          int `json:"type"`
		Order         int `json:"order,omitempty"`
		IsGraduated   int `json:"is_graduated,omitempty"`
		OpenGroupChat int `json:"open_group_chat,omitempty"`
	} `json:"departments"`
}
