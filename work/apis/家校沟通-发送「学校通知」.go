package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqSendMessage 发送「学校通知」请求
// 文档：https://developer.work.weixin.qq.com/document/path/96720#发送「学校通知」
type ReqSendMessage struct {
	RecvScope       int      `json:"recv_scope"`
	ToParentUserid  []string `json:"to_parent_userid"`
	ToStudentUserid []string `json:"to_student_userid"`
	ToParty         []string `json:"to_party"`
}

var _ bodyer = ReqSendMessage{}

func (x ReqSendMessage) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespSendMessage 发送「学校通知」响应
// 文档：https://developer.work.weixin.qq.com/document/path/96720#发送「学校通知」
type RespSendMessage struct {
	CommonResp
	InvalidParentUserid  []string `json:"invalid_parent_userid"`
	InvalidStudentUserid []string `json:"invalid_student_userid"`
	InvalidParty         []string `json:"invalid_party"`
}

var _ bodyer = RespSendMessage{}

func (x RespSendMessage) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecSendMessageText 发送「学校通知」
// 文档：https://developer.work.weixin.qq.com/document/path/96720#发送「学校通知」
func (c *ApiClient) ExecSendMessageText(req ReqSendMessageText) (RespSendMessage, error) {
	var resp RespSendMessage
	err := c.executeWXApiPost("/cgi-bin/externalcontact/message/send", req, &resp, true)
	if err != nil {
		return RespSendMessage{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespSendMessage{}, bizErr
	}
	return resp, nil
}

type ReqSendMessageText struct {
	ReqSendMessage
	Toall   int    `json:"toall"`
	Msgtype string `json:"msgtype"`
	Agentid int    `json:"agentid"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
	EnableIdTrans          int `json:"enable_id_trans"`
	EnableDuplicateCheck   int `json:"enable_duplicate_check"`
	DuplicateCheckInterval int `json:"duplicate_check_interval"`
}

// ExecSendMessageImage 发送「学校通知」
// 文档：https://developer.work.weixin.qq.com/document/path/96720#发送「学校通知」
func (c *ApiClient) ExecSendMessageImage(req ReqSendMessageImage) (RespSendMessage, error) {
	var resp RespSendMessage
	err := c.executeWXApiPost("/cgi-bin/externalcontact/message/send", req, &resp, true)
	if err != nil {
		return RespSendMessage{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespSendMessage{}, bizErr
	}
	return resp, nil
}

type ReqSendMessageImage struct {
	ReqSendMessage
	Toall   int    `json:"toall"`
	Msgtype string `json:"msgtype"`
	Agentid int    `json:"agentid"`
	Image   struct {
		MediaId string `json:"media_id"`
	} `json:"image"`
	EnableDuplicateCheck   int `json:"enable_duplicate_check"`
	DuplicateCheckInterval int `json:"duplicate_check_interval"`
}

// ExecSendMessageVoice 发送「学校通知」
// 文档：https://developer.work.weixin.qq.com/document/path/96720#发送「学校通知」
func (c *ApiClient) ExecSendMessageVoice(req ReqSendMessageVoice) (RespSendMessage, error) {
	var resp RespSendMessage
	err := c.executeWXApiPost("/cgi-bin/externalcontact/message/send", req, &resp, true)
	if err != nil {
		return RespSendMessage{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespSendMessage{}, bizErr
	}
	return resp, nil
}

type ReqSendMessageVoice struct {
	ReqSendMessage
	Toall   int    `json:"toall"`
	Msgtype string `json:"msgtype"`
	Agentid int    `json:"agentid"`
	Voice   struct {
		MediaId string `json:"media_id"`
	} `json:"voice"`
	EnableDuplicateCheck   int `json:"enable_duplicate_check"`
	DuplicateCheckInterval int `json:"duplicate_check_interval"`
}

// ExecSendMessageVideo 发送「学校通知」
// 文档：https://developer.work.weixin.qq.com/document/path/96720#发送「学校通知」
func (c *ApiClient) ExecSendMessageVideo(req ReqSendMessageVideo) (RespSendMessage, error) {
	var resp RespSendMessage
	err := c.executeWXApiPost("/cgi-bin/externalcontact/message/send", req, &resp, true)
	if err != nil {
		return RespSendMessage{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespSendMessage{}, bizErr
	}
	return resp, nil
}

type ReqSendMessageVideo struct {
	ReqSendMessage
	Toall   int    `json:"toall"`
	Msgtype string `json:"msgtype"`
	Agentid int    `json:"agentid"`
	Video   struct {
		MediaId     string `json:"media_id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"video"`
	EnableDuplicateCheck   int `json:"enable_duplicate_check"`
	DuplicateCheckInterval int `json:"duplicate_check_interval"`
}

// ExecSendMessageFile 发送「学校通知」
// 文档：https://developer.work.weixin.qq.com/document/path/96720#发送「学校通知」
func (c *ApiClient) ExecSendMessageFile(req ReqSendMessageFile) (RespSendMessage, error) {
	var resp RespSendMessage
	err := c.executeWXApiPost("/cgi-bin/externalcontact/message/send", req, &resp, true)
	if err != nil {
		return RespSendMessage{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespSendMessage{}, bizErr
	}
	return resp, nil
}

type ReqSendMessageFile struct {
	ReqSendMessage
	Toall   int    `json:"toall"`
	Msgtype string `json:"msgtype"`
	Agentid int    `json:"agentid"`
	Video   struct {
		MediaId     string `json:"media_id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"video"`
	EnableDuplicateCheck   int `json:"enable_duplicate_check"`
	DuplicateCheckInterval int `json:"duplicate_check_interval"`
}

// ExecSendMessageNews 发送「学校通知」
// 文档：https://developer.work.weixin.qq.com/document/path/96720#发送「学校通知」
func (c *ApiClient) ExecSendMessageNews(req ReqSendMessageNews) (RespSendMessage, error) {
	var resp RespSendMessage
	err := c.executeWXApiPost("/cgi-bin/externalcontact/message/send", req, &resp, true)
	if err != nil {
		return RespSendMessage{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespSendMessage{}, bizErr
	}
	return resp, nil
}

type ReqSendMessageNews struct {
	ReqSendMessage
	Toall   int    `json:"toall"`
	Msgtype string `json:"msgtype"`
	Agentid int    `json:"agentid"`
	News    struct {
		Articles []struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			Url         string `json:"url"`
			Picurl      string `json:"picurl"`
		} `json:"articles"`
	} `json:"news"`
	EnableIdTrans          int `json:"enable_id_trans"`
	EnableDuplicateCheck   int `json:"enable_duplicate_check"`
	DuplicateCheckInterval int `json:"duplicate_check_interval"`
}

// ExecSendMessageMPNews 发送「学校通知」
// 文档：https://developer.work.weixin.qq.com/document/path/96720#发送「学校通知」
func (c *ApiClient) ExecSendMessageMPNews(req ReqSendMessageMPNews) (RespSendMessage, error) {
	var resp RespSendMessage
	err := c.executeWXApiPost("/cgi-bin/externalcontact/message/send", req, &resp, true)
	if err != nil {
		return RespSendMessage{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespSendMessage{}, bizErr
	}
	return resp, nil
}

type ReqSendMessageMPNews struct {
	ReqSendMessage
	Toall   int    `json:"toall"`
	Msgtype string `json:"msgtype"`
	Agentid int    `json:"agentid"`
	Mpnews  struct {
		Articles []struct {
			Title            string `json:"title"`
			ThumbMediaId     string `json:"thumb_media_id"`
			Author           string `json:"author"`
			ContentSourceUrl string `json:"content_source_url"`
			Content          string `json:"content"`
			Digest           string `json:"digest"`
		} `json:"articles"`
	} `json:"mpnews"`
	EnableIdTrans          int `json:"enable_id_trans"`
	EnableDuplicateCheck   int `json:"enable_duplicate_check"`
	DuplicateCheckInterval int `json:"duplicate_check_interval"`
}

// ExecSendMessageProgram 发送「学校通知」
// 文档：https://developer.work.weixin.qq.com/document/path/96720#发送「学校通知」
func (c *ApiClient) ExecSendMessageProgram(req ReqSendMessageProgram) (RespSendMessage, error) {
	var resp RespSendMessage
	err := c.executeWXApiPost("/cgi-bin/externalcontact/message/send", req, &resp, true)
	if err != nil {
		return RespSendMessage{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespSendMessage{}, bizErr
	}
	return resp, nil
}

type ReqSendMessageProgram struct {
	ReqSendMessage
	Toall       int    `json:"toall"`
	Agentid     int    `json:"agentid"`
	Msgtype     string `json:"msgtype"`
	Miniprogram struct {
		Appid        string `json:"appid"`
		Title        string `json:"title"`
		ThumbMediaId string `json:"thumb_media_id"`
		Pagepath     string `json:"pagepath"`
	} `json:"miniprogram"`
	EnableIdTrans          int `json:"enable_id_trans"`
	EnableDuplicateCheck   int `json:"enable_duplicate_check"`
	DuplicateCheckInterval int `json:"duplicate_check_interval"`
}
