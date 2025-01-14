package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqBatchUpdateParentUser 批量更新家长请求
// 文档：https://developer.work.weixin.qq.com/document/path/100156#批量更新家长
type ReqBatchUpdateParentUser struct {
	Parents []*BatchUpdateParent `json:"parents"`
}

type BatchUpdateParent struct {
	ParentUserid    string                       `json:"parent_userid"`
	NewParentUserid string                       `json:"new_parent_userid,omitempty"`
	Mobile          string                       `json:"mobile"`
	Children        []*BatchUpdateParentChildren `json:"children"`
}

type BatchUpdateParentChildren struct {
	StudentUserid string `json:"student_userid"`
	Relation      string `json:"relation"`
}

var _ bodyer = ReqBatchUpdateParentUser{}

func (x ReqBatchUpdateParentUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespBatchUpdateParentUser 批量更新家长响应
// 文档：https://developer.work.weixin.qq.com/document/path/100156#批量更新家长
type RespBatchUpdateParentUser struct {
	CommonResp
	ResultList []struct {
		ParentUserid string `json:"parent_userid"`
		Errcode      int    `json:"errcode"`
		Errmsg       string `json:"errmsg"`
	} `json:"result_list"`
}

var _ bodyer = RespBatchUpdateParentUser{}

func (x RespBatchUpdateParentUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecBatchUpdateParentUser 批量更新家长
// 文档：https://developer.work.weixin.qq.com/document/path/100156#批量更新家长
func (c *ApiClient) ExecBatchUpdateParentUser(req ReqBatchUpdateParentUser) (RespBatchUpdateParentUser, error) {
	var resp RespBatchUpdateParentUser
	err := c.executeWXApiPost("/cgi-bin/school/user/batch_update_parent", req, &resp, true)
	if err != nil {
		return RespBatchUpdateParentUser{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespBatchUpdateParentUser{}, bizErr
	}
	return resp, nil
}
