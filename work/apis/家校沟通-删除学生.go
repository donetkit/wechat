package apis

import (
	"encoding/json"
	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqDeleteStudentUser 删除学生请求
// 文档：https://developer.work.weixin.qq.com/document/path/100146#删除学生
type ReqDeleteStudentUser struct {
	Userid string `json:"userid"`
}

var _ urlValuer = ReqDeleteStudentUser{}

func (x ReqDeleteStudentUser) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespDeleteStudentUser 删除学生响应
// 文档：https://developer.work.weixin.qq.com/document/path/100146#删除学生
type RespDeleteStudentUser struct {
	CommonResp
}

var _ bodyer = RespDeleteStudentUser{}

func (x RespDeleteStudentUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecDeleteStudentUser 删除学生
// 文档：https://developer.work.weixin.qq.com/document/path/100146#删除学生
func (c *ApiClient) ExecDeleteStudentUser(req ReqDeleteStudentUser) (RespDeleteStudentUser, error) {
	var resp RespDeleteStudentUser
	err := c.executeWXApiGet("/cgi-bin/school/user/delete_student", req, &resp, true)
	if err != nil {
		return RespDeleteStudentUser{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespDeleteStudentUser{}, bizErr
	}
	return resp, nil
}
