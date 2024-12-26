package apis

import (
	"encoding/json"

	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetcorpconfVacation 获取企业假期管理配置请求
// 文档：https://developer.work.weixin.qq.com/document/path/94211#获取企业假期管理配置
type ReqGetcorpconfVacation struct {
}

var _ urlValuer = ReqGetcorpconfVacation{}

func (x ReqGetcorpconfVacation) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespGetcorpconfVacation 获取企业假期管理配置响应
// 文档：https://developer.work.weixin.qq.com/document/path/94211#获取企业假期管理配置
type RespGetcorpconfVacation struct {
	CommonResp
	Lists []struct {
		Id           int    `json:"id"`
		Name         string `json:"name"`
		TimeAttr     int    `json:"time_attr"`
		DurationType int    `json:"duration_type"`
		QuotaAttr    struct {
			Type              int  `json:"type"`
			AutoresetTime     int  `json:"autoreset_time"`
			AutoresetDuration int  `json:"autoreset_duration"`
			QuotaRuleType     int  `json:"quota_rule_type"`
			AtEntryDate       bool `json:"at_entry_date"`
			AutoResetMonthDay int  `json:"auto_reset_month_day"`
		} `json:"quota_attr"`
		PerdayDuration     int `json:"perday_duration"`
		IsNewovertime      int `json:"is_newovertime"`
		EnterCompTimeLimit int `json:"enter_comp_time_limit"`
		ExpireRule         struct {
			Type     int `json:"type"`
			Duration int `json:"duration"`
			Date     struct {
				Month int `json:"month"`
				Day   int `json:"day"`
			} `json:"date"`
			ExternDurationEnable bool `json:"extern_duration_enable"`
			ExternDuration       struct {
				Month int `json:"month"`
				Day   int `json:"day"`
			} `json:"extern_duration"`
		} `json:"expire_rule"`
	} `json:"lists"`
}

var _ bodyer = RespGetcorpconfVacation{}

func (x RespGetcorpconfVacation) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetcorpconfVacation 获取企业假期管理配置
// 文档：https://developer.work.weixin.qq.com/document/path/94211#获取企业假期管理配置
func (c *ApiClient) ExecGetcorpconfVacation(req ReqGetcorpconfVacation) (RespGetcorpconfVacation, error) {
	var resp RespGetcorpconfVacation
	err := c.executeWXApiGet("/cgi-bin/oa/vacation/getcorpconf", req, &resp, true)
	if err != nil {
		return RespGetcorpconfVacation{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetcorpconfVacation{}, bizErr
	}
	return resp, nil
}
