package apis

import (
	"encoding/json"
	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqListUserSchool 获取部门学生详情请求
// 文档：https://developer.work.weixin.qq.com/document/path/96739#获取部门学生详情
type ReqListUserSchool struct {
	DepartmentID int `json:"department_id"`
}

var _ urlValuer = ReqListUserSchool{}

func (x ReqListUserSchool) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespListUserSchool 获取部门学生详情响应
// 文档：https://developer.work.weixin.qq.com/document/path/96739#获取部门学生详情
type RespListUserSchool struct {
	CommonResp
	Students []struct {
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
	} `json:"students"`
}

var _ bodyer = RespListUserSchool{}

func (x RespListUserSchool) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecListUserSchool 获取部门学生详情
// 文档：https://developer.work.weixin.qq.com/document/path/96739#获取部门学生详情
func (c *ApiClient) ExecListUserSchool(req ReqListUserSchool) (RespListUserSchool, error) {
	var resp RespListUserSchool
	err := c.executeWXApiGet("/cgi-bin/school/user/list", req, &resp, true)
	if err != nil {
		return RespListUserSchool{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespListUserSchool{}, bizErr
	}
	return resp, nil
}
