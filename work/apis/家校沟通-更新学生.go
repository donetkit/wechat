package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqUpdateStudentUser 更新学生请求
// 文档：https://developer.work.weixin.qq.com/document/path/100147#更新学生
type ReqUpdateStudentUser struct {
	StudentUserId    string `json:"student_userid"`
	Mobile           string `json:"mobile"`
	NewStudentUserid string `json:"new_student_userid,omitempty"`
	Name             string `json:"name"`
	Department       []int  `json:"department"`
}

var _ bodyer = ReqUpdateStudentUser{}

func (x ReqUpdateStudentUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespUpdateStudentUser 更新学生响应
// 文档：https://developer.work.weixin.qq.com/document/path/100147#更新学生
type RespUpdateStudentUser struct {
	CommonResp
}

var _ bodyer = RespUpdateStudentUser{}

func (x RespUpdateStudentUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecUpdateStudentUser 更新学生
// 文档：https://developer.work.weixin.qq.com/document/path/100147#更新学生
func (c *ApiClient) ExecUpdateStudentUser(req ReqUpdateStudentUser) (RespUpdateStudentUser, error) {
	var resp RespUpdateStudentUser
	err := c.executeWXApiPost("/cgi-bin/school/user/update_student", req, &resp, true)
	if err != nil {
		return RespUpdateStudentUser{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespUpdateStudentUser{}, bizErr
	}
	return resp, nil
}
