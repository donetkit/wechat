package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqCreateStudentUser 创建学生请求
// 文档：https://developer.work.weixin.qq.com/document/path/100145#创建学生
type ReqCreateStudentUser struct {
	StudentUserId string `json:"student_userid"`
	Mobile        string `json:"mobile"`
	ToInvite      bool   `json:"to_invite"`
	Name          string `json:"name"`
	Department    []int  `json:"department"`
}

var _ bodyer = ReqCreateStudentUser{}

func (x ReqCreateStudentUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespCreateStudentUser 创建学生响应
// 文档：https://developer.work.weixin.qq.com/document/path/100145#创建学生
type RespCreateStudentUser struct {
	CommonResp
}

var _ bodyer = RespCreateStudentUser{}

func (x RespCreateStudentUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecCreateStudentUser 创建学生
// 文档：https://developer.work.weixin.qq.com/document/path/100145#创建学生
func (c *ApiClient) ExecCreateStudentUser(req ReqCreateStudentUser) (RespCreateStudentUser, error) {
	var resp RespCreateStudentUser
	err := c.executeWXApiPost("/cgi-bin/school/user/create_student", req, &resp, true)
	if err != nil {
		return RespCreateStudentUser{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespCreateStudentUser{}, bizErr
	}
	return resp, nil
}
