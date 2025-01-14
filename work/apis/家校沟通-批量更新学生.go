package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqBatchUpdateStudentUser 批量更新学生请求
// 文档：https://developer.work.weixin.qq.com/document/path/100150#批量更新学生
type ReqBatchUpdateStudentUser struct {
	Students []*BatchUpdateStudent `json:"students"`
}

type BatchUpdateStudent struct {
	StudentUserId    string `json:"student_userid"`
	Mobile           string `json:"mobile"`
	NewStudentUserId string `json:"new_student_userid,omitempty"`
	Name             string `json:"name"`
	Department       []int  `json:"department"`
}

var _ bodyer = ReqBatchUpdateStudentUser{}

func (x ReqBatchUpdateStudentUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespBatchUpdateStudentUser 批量更新学生响应
// 文档：https://developer.work.weixin.qq.com/document/path/100150#批量更新学生
type RespBatchUpdateStudentUser struct {
	CommonResp
	ResultList []*ResultStudentUserList `json:"result_list"`
}

var _ bodyer = RespBatchUpdateStudentUser{}

func (x RespBatchUpdateStudentUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecBatchUpdateStudentUser 批量更新学生
// 文档：https://developer.work.weixin.qq.com/document/path/100150#批量更新学生
func (c *ApiClient) ExecBatchUpdateStudentUser(req ReqBatchUpdateStudentUser) (RespBatchUpdateStudentUser, error) {
	var resp RespBatchUpdateStudentUser
	err := c.executeWXApiPost("/cgi-bin/school/user/batch_update_student", req, &resp, true)
	if err != nil {
		return RespBatchUpdateStudentUser{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespBatchUpdateStudentUser{}, bizErr
	}
	return resp, nil
}
