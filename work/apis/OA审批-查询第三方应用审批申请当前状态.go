package apis

import (
	"encoding/json"
)

type ReqGetOpenApprovalDataOa struct {
	ThirdNo string `json:"thirdNo"`
}

var _ bodyer = ReqGetOpenApprovalDataOa{}

func (x ReqGetOpenApprovalDataOa) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespGetOpenApprovalDataOa struct {
	CommonResp
	Data struct {
		ThirdNo        string `json:"ThirdNo"`
		OpenTemplateId string `json:"OpenTemplateId"`
		OpenSpName     string `json:"OpenSpName"`
		OpenSpstatus   int    `json:"OpenSpstatus"`
		ApplyTime      int    `json:"ApplyTime"`
		ApplyUsername  string `json:"ApplyUsername"`
		ApplyUserParty string `json:"ApplyUserParty"`
		ApplyUserImage string `json:"ApplyUserImage"`
		ApplyUserId    string `json:"ApplyUserId"`
		ApprovalNodes  struct {
			ApprovalNode []struct {
				NodeStatus int `json:"NodeStatus"`
				NodeAttr   int `json:"NodeAttr"`
				NodeType   int `json:"NodeType"`
				Items      struct {
					Item []struct {
						ItemName   string `json:"ItemName"`
						ItemParty  string `json:"ItemParty"`
						ItemImage  string `json:"ItemImage"`
						ItemUserId string `json:"ItemUserId"`
						ItemStatus int    `json:"ItemStatus"`
						ItemSpeech string `json:"ItemSpeech"`
						ItemOpTime int    `json:"ItemOpTime"`
					} `json:"Item"`
				} `json:"Items"`
			} `json:"ApprovalNode"`
		} `json:"ApprovalNodes"`
		NotifyNodes struct {
			NotifyNode []struct {
				ItemName   string `json:"ItemName"`
				ItemParty  string `json:"ItemParty"`
				ItemImage  string `json:"ItemImage"`
				ItemUserId string `json:"ItemUserId"`
			} `json:"NotifyNode"`
		} `json:"NotifyNodes"`
		Approverstep int `json:"approverstep"`
	} `json:"data"`
}

var _ bodyer = RespGetOpenApprovalDataOa{}

func (x RespGetOpenApprovalDataOa) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecGetOpenApprovalDataOa(req ReqGetOpenApprovalDataOa) (RespGetOpenApprovalDataOa, error) {
	var resp RespGetOpenApprovalDataOa
	err := c.executeWXApiPost("/cgi-bin/corp/getopenapprovaldata", req, &resp, true)
	if err != nil {
		return RespGetOpenApprovalDataOa{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetOpenApprovalDataOa{}, bizErr
	}
	return resp, nil
}
