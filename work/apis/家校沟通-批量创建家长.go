package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqBatchCreateParentUser 批量创建家长请求
// 文档：https://developer.work.weixin.qq.com/document/path/100154#批量创建家长
type ReqBatchCreateParentUser struct {
	Parents []*BatchCreateParent `json:"parents"`
}

type BatchCreateParent struct {
	ParentUserid string                       `json:"parent_userid"`
	Mobile       string                       `json:"mobile"`
	ToInvite     bool                         `json:"to_invite,omitempty"`
	Children     []*BatchCreateParentChildren `json:"children"`
}

type BatchCreateParentChildren struct {
	StudentUserid string `json:"student_userid"`
	Relation      string `json:"relation"`
}

var _ bodyer = ReqBatchCreateParentUser{}

func (x ReqBatchCreateParentUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespBatchCreateParentUser 批量创建家长响应
// 文档：https://developer.work.weixin.qq.com/document/path/100154#批量创建家长
type RespBatchCreateParentUser struct {
	CommonResp
	ResultList []*ResultBatchCreateParentList `json:"result_list"`
}

type ResultBatchCreateParentList struct {
	ParentUserid string `json:"parent_userid"`
	CommonResp
}

var _ bodyer = RespBatchCreateParentUser{}

func (x RespBatchCreateParentUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecBatchCreateParentUser 批量创建家长
// 文档：https://developer.work.weixin.qq.com/document/path/100154#批量创建家长
func (c *ApiClient) ExecBatchCreateParentUser(req ReqBatchCreateParentUser) (RespBatchCreateParentUser, error) {
	var resp RespBatchCreateParentUser
	err := c.executeWXApiPost("/cgi-bin/school/user/batch_create_parent", req, &resp, true)
	if err != nil {
		return RespBatchCreateParentUser{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespBatchCreateParentUser{}, bizErr
	}
	return resp, nil
}
