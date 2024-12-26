package apis

import (
	"encoding/json"
)

type (
	ReqUpdateUser struct {
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
)

var _ bodyer = ReqUpdateUser{}

func (x ReqUpdateUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespUpdateUser struct {
	CommonResp
}

var _ bodyer = RespUpdateUser{}

func (x RespUpdateUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecUpdateUser(req ReqUpdateUser) (RespUpdateUser, error) {
	var resp RespUpdateUser
	err := c.executeWXApiPost("/cgi-bin/user/update", req, &resp, true)
	if err != nil {
		return RespUpdateUser{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespUpdateUser{}, bizErr
	}
	return resp, nil
}
