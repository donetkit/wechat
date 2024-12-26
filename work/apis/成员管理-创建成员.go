package apis

import (
	"encoding/json"
)

type (
	ReqCreateUser struct {
		Userid           string                     `json:"userid"`
		Name             string                     `json:"name,omitempty"`
		Alias            string                     `json:"alias,omitempty"`
		Mobile           string                     `json:"mobile,omitempty"`
		Department       []int                      `json:"department,omitempty"`
		Order            []int                      `json:"order,omitempty"`
		Position         string                     `json:"position,omitempty"`
		Gender           string                     `json:"gender,omitempty"`
		Email            string                     `json:"email,omitempty"`
		BizMail          string                     `json:"biz_mail,omitempty"`
		IsLeaderInDept   []int                      `json:"is_leader_in_dept,omitempty"`
		DirectLeader     []string                   `json:"direct_leader,omitempty"`
		Enable           int                        `json:"enable,omitempty"`
		AvatarMediaid    string                     `json:"avatar_mediaid,omitempty"`
		Telephone        string                     `json:"telephone,omitempty"`
		Address          string                     `json:"address,omitempty"`
		MainDepartment   int                        `json:"main_department,omitempty"`
		Extattr          *CreateUserExtattr         `json:"extattr,omitempty"`
		ToInvite         bool                       `json:"to_invite,omitempty"`
		ExternalPosition string                     `json:"external_position,omitempty"`
		ExternalProfile  *CreateUserExternalProfile `json:"external_profile,omitempty"`
	}

	CreateUserExtattr struct {
		Attrs []CreateUserExtattrAttrs `json:"attrs"`
	}

	CreateUserExtattrAttrs struct {
		Type int                         `json:"type"`
		Name string                      `json:"name"`
		Text *CreateUserExtattrAttrsText `json:"text,omitempty"`
		Web  *CreateUserExtattrAttrsWeb  `json:"web,omitempty"`
	}

	CreateUserExtattrAttrsText struct {
		Value string `json:"value"`
	}

	CreateUserExtattrAttrsWeb struct {
		Url   string `json:"url"`
		Title string `json:"title"`
	}

	CreateUserExternalProfile struct {
		ExternalCorpName string                                   `json:"external_corp_name"`
		WechatChannels   *CreateUserExternalProfileWechatChannels `json:"wechat_channels,omitempty"`
		ExternalAttr     []struct {
			CreateUserExtattrAttrs
			Miniprogram *CreateUserExtattrAttrsTextMiniprogram `json:"miniprogram,omitempty"`
		} `json:"external_attr"`
	}

	CreateUserExternalProfileWechatChannels struct {
		Nickname string `json:"nickname"`
	}

	CreateUserExtattrAttrsTextMiniprogram struct {
		Appid    string `json:"appid"`
		Pagepath string `json:"pagepath"`
		Title    string `json:"title"`
	}
)

var _ bodyer = ReqCreateUser{}

func (x ReqCreateUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespCreateUser struct {
	CommonResp
	CreatedDepartmentList struct {
		DepartmentInfo []struct {
			Name string `json:"name"`
			Id   int    `json:"id"`
		} `json:"department_info"`
	} `json:"created_department_list"`
}

var _ bodyer = RespCreateUser{}

func (x RespCreateUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecCreateUser(req ReqCreateUser) (RespCreateUser, error) {
	var resp RespCreateUser
	err := c.executeWXApiPost("/cgi-bin/user/create", req, &resp, true)
	if err != nil {
		return RespCreateUser{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespCreateUser{}, bizErr
	}
	return resp, nil
}
