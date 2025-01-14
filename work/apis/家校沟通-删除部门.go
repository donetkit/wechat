package apis

import (
	"encoding/json"
	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqDeleteDepartmentSchool 删除部门请求
// 文档：https://developer.work.weixin.qq.com/document/path/100160#删除部门
type ReqDeleteDepartmentSchool struct {
	ID int `json:"id"`
}

var _ urlValuer = ReqDeleteDepartmentSchool{}

func (x ReqDeleteDepartmentSchool) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespDeleteDepartmentSchool 删除部门响应
// 文档：https://developer.work.weixin.qq.com/document/path/100160#删除部门
type RespDeleteDepartmentSchool struct {
	CommonResp
}

var _ bodyer = RespDeleteDepartmentSchool{}

func (x RespDeleteDepartmentSchool) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecDeleteDepartmentSchool 删除部门
// 文档：https://developer.work.weixin.qq.com/document/path/100160#删除部门
func (c *ApiClient) ExecDeleteDepartmentSchool(req ReqDeleteDepartmentSchool) (RespDeleteDepartmentSchool, error) {
	var resp RespDeleteDepartmentSchool
	err := c.executeWXApiGet("/cgi-bin/school/department/delete", req, &resp, true)
	if err != nil {
		return RespDeleteDepartmentSchool{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespDeleteDepartmentSchool{}, bizErr
	}
	return resp, nil
}
