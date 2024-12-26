package apis

import (
	"encoding/json"
)

type ReqGetCheckInData struct {
	Opencheckindatatype int      `json:"opencheckindatatype"`
	Starttime           int      `json:"starttime"`
	Endtime             int      `json:"endtime"`
	Useridlist          []string `json:"useridlist"`
}

var _ bodyer = ReqGetCheckInData{}

func (x ReqGetCheckInData) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespGetCheckInData struct {
	CommonResp
	Checkindata []struct {
		Userid         string   `json:"userid"`
		Groupname      string   `json:"groupname"`
		CheckinType    string   `json:"checkin_type"`
		ExceptionType  string   `json:"exception_type"`
		CheckinTime    int      `json:"checkin_time"`
		LocationTitle  string   `json:"location_title"`
		LocationDetail string   `json:"location_detail"`
		Wifiname       string   `json:"wifiname"`
		Notes          string   `json:"notes"`
		Wifimac        string   `json:"wifimac"`
		Mediaids       []string `json:"mediaids"`
		SchCheckinTime int      `json:"sch_checkin_time"`
		Groupid        int      `json:"groupid"`
		ScheduleId     int      `json:"schedule_id"`
		TimelineId     int      `json:"timeline_id"`
		Lat            int      `json:"lat,omitempty"`
		Lng            int      `json:"lng,omitempty"`
		Deviceid       string   `json:"deviceid,omitempty"`
	} `json:"checkindata"`
}

var _ bodyer = RespGetCheckInData{}

func (x RespGetCheckInData) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecGetCheckInData(req ReqGetCheckInData) (RespGetCheckInData, error) {
	var resp RespGetCheckInData
	err := c.executeWXApiPost("/cgi-bin/checkin/getcheckindata", req, &resp, true)
	if err != nil {
		return RespGetCheckInData{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetCheckInData{}, bizErr
	}
	return resp, nil
}
