package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqUpdateParentUser 更新家长请求
// 文档：https://developer.work.weixin.qq.com/document/path/100153#更新家长
type ReqUpdateParentUser struct {
	ParentUserid    string                `json:"parent_userid"`
	NewParentUserid string                `json:"new_parent_userid"`
	Mobile          string                `json:"mobile"`
	Children        []*ChildrenParentUser `json:"children"`
}

type ChildrenParentUser struct {
	StudentUserid string `json:"student_userid"`
	Relation      string `json:"relation"`
}

var _ bodyer = ReqUpdateParentUser{}

func (x ReqUpdateParentUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespUpdateParentUser 更新家长响应
// 文档：https://developer.work.weixin.qq.com/document/path/100153#更新家长
type RespUpdateParentUser struct {
	CommonResp
}

var _ bodyer = RespUpdateParentUser{}

func (x RespUpdateParentUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecUpdateParentUser 更新家长
// 文档：https://developer.work.weixin.qq.com/document/path/100153#更新家长
func (c *ApiClient) ExecUpdateParentUser(req ReqUpdateParentUser) (RespUpdateParentUser, error) {
	var resp RespUpdateParentUser
	err := c.executeWXApiPost("/cgi-bin/school/user/update_parent", req, &resp, true)
	if err != nil {
		return RespUpdateParentUser{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespUpdateParentUser{}, bizErr
	}
	return resp, nil
}
