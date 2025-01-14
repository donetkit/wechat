package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqBatchDeleteParentUser 批量删除家长请求
// 文档：https://developer.work.weixin.qq.com/document/path/100155#批量删除家长
type ReqBatchDeleteParentUser struct {
	UserIdList []string `json:"useridlist"`
}

var _ bodyer = ReqBatchDeleteParentUser{}

func (x ReqBatchDeleteParentUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespBatchDeleteParentUser 批量删除家长响应
// 文档：https://developer.work.weixin.qq.com/document/path/100155#批量删除家长
type RespBatchDeleteParentUser struct {
	CommonResp
	ResultList []struct {
		ParentUserid string `json:"parent_userid"`
		Errcode      int    `json:"errcode"`
		Errmsg       string `json:"errmsg"`
	} `json:"result_list"`
}

var _ bodyer = RespBatchDeleteParentUser{}

func (x RespBatchDeleteParentUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecBatchDeleteParentUser 批量删除家长
// 文档：https://developer.work.weixin.qq.com/document/path/100155#批量删除家长
func (c *ApiClient) ExecBatchDeleteParentUser(req ReqBatchDeleteParentUser) (RespBatchDeleteParentUser, error) {
	var resp RespBatchDeleteParentUser
	err := c.executeWXApiPost("/cgi-bin/school/user/batch_delete_parent", req, &resp, true)
	if err != nil {
		return RespBatchDeleteParentUser{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespBatchDeleteParentUser{}, bizErr
	}
	return resp, nil
}
