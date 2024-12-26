package apis

import (
	"encoding/json"
)

type (
	ReqGetApprovalInfoOa struct {
		Starttime string                        `json:"starttime"`
		Endtime   string                        `json:"endtime"`
		NewCursor string                        `json:"new_cursor"`
		Size      int                           `json:"size"`
		Filters   []ReqGetApprovalInfoOaFilters `json:"filters"`
	}

	ReqGetApprovalInfoOaFilters struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
)

var _ bodyer = ReqGetApprovalInfoOa{}

func (x ReqGetApprovalInfoOa) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespGetApprovalInfoOa struct {
	CommonResp
	SpNoList []string `json:"sp_no_list"`
}

var _ bodyer = RespGetApprovalInfoOa{}

func (x RespGetApprovalInfoOa) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecGetApprovalInfoOa(req ReqGetApprovalInfoOa) (RespGetApprovalInfoOa, error) {
	var resp RespGetApprovalInfoOa
	err := c.executeWXApiPost("/cgi-bin/oa/getapprovalinfo", req, &resp, true)
	if err != nil {
		return RespGetApprovalInfoOa{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetApprovalInfoOa{}, bizErr
	}
	return resp, nil
}
