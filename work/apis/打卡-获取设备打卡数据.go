package apis

import (
	"encoding/json"
)

type ReqGetHardwareCheckInData struct {
	FilterType int      `json:"filter_type"`
	Starttime  int      `json:"starttime"`
	Endtime    int      `json:"endtime"`
	Useridlist []string `json:"useridlist"`
}

var _ bodyer = ReqGetHardwareCheckInData{}

func (x ReqGetHardwareCheckInData) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespGetHardwareCheckInData struct {
	CommonResp
	Checkindata []struct {
		Userid      string `json:"userid"`
		CheckinTime int    `json:"checkin_time"`
		DeviceSn    string `json:"device_sn"`
		DeviceName  string `json:"device_name"`
	} `json:"checkindata"`
}

var _ bodyer = RespGetHardwareCheckInData{}

func (x RespGetHardwareCheckInData) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecGetHardwareCheckInData(req ReqGetHardwareCheckInData) (RespGetHardwareCheckInData, error) {
	var resp RespGetHardwareCheckInData
	err := c.executeWXApiPost("/cgi-bin/checkin/get_hardware_checkin_data", req, &resp, true)
	if err != nil {
		return RespGetHardwareCheckInData{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetHardwareCheckInData{}, bizErr
	}
	return resp, nil
}
