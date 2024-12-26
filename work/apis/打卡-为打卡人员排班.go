package apis

import (
	"encoding/json"
)

type (
	ReqSetCheckInSchedulist struct {
		Groupid   int                            `json:"groupid"`
		Items     []ReqSetCheckInSchedulistItems `json:"items"`
		Yearmonth int                            `json:"yearmonth"`
	}

	ReqSetCheckInSchedulistItems struct {
		Userid     string `json:"userid"`
		Day        int    `json:"day"`
		ScheduleId int    `json:"schedule_id"`
	}
)

var _ bodyer = ReqSetCheckInSchedulist{}

func (x ReqSetCheckInSchedulist) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespSetCheckInSchedulist struct {
	CommonResp
}

var _ bodyer = RespSetCheckInSchedulist{}

func (x RespSetCheckInSchedulist) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecSetCheckInSchedulist(req ReqSetCheckInSchedulist) (RespSetCheckInSchedulist, error) {
	var resp RespSetCheckInSchedulist
	err := c.executeWXApiPost("/cgi-bin/checkin/setcheckinschedulist", req, &resp, true)
	if err != nil {
		return RespSetCheckInSchedulist{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespSetCheckInSchedulist{}, bizErr
	}
	return resp, nil
}
