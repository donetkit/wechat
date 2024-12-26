package apis

import (
	"encoding/json"
	"net/url"
)

// ReqMiniProgramJsCode2Session code2Session
// 文档：https://developer.work.weixin.qq.com/document/path/96959#code2Session
type ReqMiniProgramJsCode2Session struct {
	JsCode    string `json:"js_code"`    // 登录时获取的 code
	GrantType string `json:"grant_type"` // 此处固定为authorization_code
}

var _ urlValuer = ReqMiniProgramJsCode2Session{}

func (x ReqMiniProgramJsCode2Session) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespMiniProgramJsCode2Session code2Session
// 文档：https://developer.work.weixin.qq.com/document/path/96959#code2Session
type RespMiniProgramJsCode2Session struct {
	Corpid     string `json:"corpid"`
	Userid     string `json:"userid"`
	SessionKey string `json:"session_key"`
	CommonResp
}

var _ bodyer = RespMiniProgramJsCode2Session{}

func (x RespMiniProgramJsCode2Session) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecMiniProgramJsCode2Session code2Session
// 文档：https://developer.work.weixin.qq.com/document/path/96959#code2Session
func (c *ApiClient) ExecMiniProgramJsCode2Session(req ReqMiniProgramJsCode2Session) (RespMiniProgramJsCode2Session, error) {
	var resp RespMiniProgramJsCode2Session
	req.GrantType = "authorization_code"
	err := c.executeWXApiGet("/cgi-bin/miniprogram/jscode2session", req, &resp, true)
	if err != nil {
		return RespMiniProgramJsCode2Session{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespMiniProgramJsCode2Session{}, bizErr
	}
	return resp, nil
}
