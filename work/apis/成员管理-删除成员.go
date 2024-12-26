package apis

import (
	"encoding/json"
	"net/url"
)

type ReqDeleteUser struct {
	Userid string `json:"userid"`
}

var _ urlValuer = ReqDeleteUser{}

func (x ReqDeleteUser) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

type RespDeleteUser struct {
	CommonResp
}

var _ bodyer = RespDeleteUser{}

func (x RespDeleteUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecDeleteUser(req ReqDeleteUser) (RespDeleteUser, error) {
	var resp RespDeleteUser
	err := c.executeWXApiGet("/cgi-bin/user/delete", req, &resp, true)
	if err != nil {
		return RespDeleteUser{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespDeleteUser{}, bizErr
	}
	return resp, nil
}
