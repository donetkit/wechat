package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqCreateParentUser 创建家长请求
// 文档：https://developer.work.weixin.qq.com/document/path/100151#创建家长
type ReqCreateParentUser struct {
	ParentUserid string            `json:"parent_userid"`
	Mobile       string            `json:"mobile"`
	ToInvite     bool              `json:"to_invite"`
	Children     []*ChildrenParent `json:"children"`
}

type ChildrenParent struct {
	StudentUserId string `json:"student_userid"`
	Relation      string `json:"relation"`
}

var _ bodyer = ReqCreateParentUser{}

func (x ReqCreateParentUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespCreateParentUser 创建家长响应
// 文档：https://developer.work.weixin.qq.com/document/path/100151#创建家长
type RespCreateParentUser struct {
	CommonResp
}

var _ bodyer = RespCreateParentUser{}

func (x RespCreateParentUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecCreateParentUser 创建家长
// 文档：https://developer.work.weixin.qq.com/document/path/100151#创建家长
func (c *ApiClient) ExecCreateParentUser(req ReqCreateParentUser) (RespCreateParentUser, error) {
	var resp RespCreateParentUser
	err := c.executeWXApiPost("/cgi-bin/school/user/create_parent", req, &resp, true)
	if err != nil {
		return RespCreateParentUser{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespCreateParentUser{}, bizErr
	}
	return resp, nil
}
