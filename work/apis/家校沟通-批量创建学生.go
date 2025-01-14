package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqBatchCreateStudentUser 批量创建学生请求
// 文档：https://developer.work.weixin.qq.com/document/path/100148#批量创建学生
type ReqBatchCreateStudentUser struct {
	Students []*Student `json:"students"`
}

type Student struct {
	StudentUserId string `json:"student_userid"`
	Mobile        string `json:"mobile"`
	ToInvite      bool   `json:"to_invite,omitempty"`
	Name          string `json:"name"`
	Department    []int  `json:"department"`
}

var _ bodyer = ReqBatchCreateStudentUser{}

func (x ReqBatchCreateStudentUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespBatchCreateStudentUser 批量创建学生响应
// 文档：https://developer.work.weixin.qq.com/document/path/100148#批量创建学生
type RespBatchCreateStudentUser struct {
	CommonResp
	ResultList []*ResultStudentUserList `json:"result_list"`
}

type ResultStudentUserList struct {
	StudentUserid string `json:"student_userid"`
	CommonResp
}

var _ bodyer = RespBatchCreateStudentUser{}

func (x RespBatchCreateStudentUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecBatchCreateStudentUser 批量创建学生
// 文档：https://developer.work.weixin.qq.com/document/path/100148#批量创建学生
func (c *ApiClient) ExecBatchCreateStudentUser(req ReqBatchCreateStudentUser) (RespBatchCreateStudentUser, error) {
	var resp RespBatchCreateStudentUser
	err := c.executeWXApiPost("/cgi-bin/school/user/batch_create_student", req, &resp, true)
	if err != nil {
		return RespBatchCreateStudentUser{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespBatchCreateStudentUser{}, bizErr
	}
	return resp, nil
}
