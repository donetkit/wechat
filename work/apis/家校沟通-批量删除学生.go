package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqBatchDeleteStudentUser 批量删除学生请求
// 文档：https://developer.work.weixin.qq.com/document/path/100149#批量删除学生
type ReqBatchDeleteStudentUser struct {
	UserIdList []string `json:"useridlist"`
}

var _ bodyer = ReqBatchDeleteStudentUser{}

func (x ReqBatchDeleteStudentUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespBatchDeleteStudentUser 批量删除学生响应
// 文档：https://developer.work.weixin.qq.com/document/path/100149#批量删除学生
type RespBatchDeleteStudentUser struct {
	CommonResp
	ResultList []*ResultStudentUserList `json:"result_list"`
}

var _ bodyer = RespBatchDeleteStudentUser{}

func (x RespBatchDeleteStudentUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecBatchDeleteStudentUser 批量删除学生
// 文档：https://developer.work.weixin.qq.com/document/path/100149#批量删除学生
func (c *ApiClient) ExecBatchDeleteStudentUser(req ReqBatchDeleteStudentUser) (RespBatchDeleteStudentUser, error) {
	var resp RespBatchDeleteStudentUser
	err := c.executeWXApiPost("/cgi-bin/school/user/batch_delete_student", req, &resp, true)
	if err != nil {
		return RespBatchDeleteStudentUser{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespBatchDeleteStudentUser{}, bizErr
	}
	return resp, nil
}
