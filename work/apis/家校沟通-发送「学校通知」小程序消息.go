package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqSendMessageMiniProgram 发送「学校通知」小程序消息请求
// 文档：https://developer.work.weixin.qq.com/document/path/96723#发送「学校通知」小程序消息
type ReqSendMessageMiniProgram struct {
	ToExternalUser  []string `json:"to_external_user"`
	ToParentUserid  []string `json:"to_parent_userid"`
	ToStudentUserid []string `json:"to_student_userid"`
	ToParty         []string `json:"to_party"`
	Toall           int      `json:"toall"`
	Agentid         int      `json:"agentid"`
	Msgtype         string   `json:"msgtype"`
	Miniprogram     struct {
		Appid        string `json:"appid"`
		Title        string `json:"title"`
		ThumbMediaId string `json:"thumb_media_id"`
		Pagepath     string `json:"pagepath"`
	} `json:"miniprogram"`
	EnableIdTrans int `json:"enable_id_trans"`
}

var _ bodyer = ReqSendMessageMiniProgram{}

func (x ReqSendMessageMiniProgram) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespSendMessageMiniProgram 发送「学校通知」小程序消息响应
// 文档：https://developer.work.weixin.qq.com/document/path/96723#发送「学校通知」小程序消息
type RespSendMessageMiniProgram struct {
	CommonResp
	InvalidExternalUser  []string `json:"invalid_external_user"`
	InvalidParentUserid  []string `json:"invalid_parent_userid"`
	InvalidStudentUserid []string `json:"invalid_student_userid"`
	InvalidParty         []string `json:"invalid_party"`
}

var _ bodyer = RespSendMessageMiniProgram{}

func (x RespSendMessageMiniProgram) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecSendMessageMiniProgram 发送「学校通知」小程序消息
// 文档：https://developer.work.weixin.qq.com/document/path/96723#发送「学校通知」小程序消息
func (c *ApiClient) ExecSendMessageMiniProgram(req ReqSendMessageMiniProgram) (RespSendMessageMiniProgram, error) {
	var resp RespSendMessageMiniProgram
	err := c.executeWXApiPost("/cgi-bin/externalcontact/message/send", req, &resp, true)
	if err != nil {
		return RespSendMessageMiniProgram{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespSendMessageMiniProgram{}, bizErr
	}
	return resp, nil
}
