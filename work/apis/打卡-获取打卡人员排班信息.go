package apis

import (
	"encoding/json"
)

type ReqGetCheckInSchedulist struct {
	Starttime  int      `json:"starttime"`
	Endtime    int      `json:"endtime"`
	Useridlist []string `json:"useridlist"`
}

var _ bodyer = ReqGetCheckInSchedulist{}

func (x ReqGetCheckInSchedulist) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespGetCheckInSchedulist struct {
	CommonResp
	ScheduleList []struct {
		Userid    string `json:"userid"`
		Yearmonth int    `json:"yearmonth"`
		Groupid   int    `json:"groupid"`
		Groupname string `json:"groupname"`
		Schedule  struct {
			ScheduleList []struct {
				Day          int `json:"day"`
				ScheduleInfo struct {
					ScheduleId   int    `json:"schedule_id"`
					ScheduleName string `json:"schedule_name"`
					TimeSection  []struct {
						Id               int `json:"id"`
						WorkSec          int `json:"work_sec"`
						OffWorkSec       int `json:"off_work_sec"`
						RemindWorkSec    int `json:"remind_work_sec"`
						RemindOffWorkSec int `json:"remind_off_work_sec"`
					} `json:"time_section"`
				} `json:"schedule_info"`
			} `json:"scheduleList"`
		} `json:"schedule"`
	} `json:"schedule_list"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

var _ bodyer = RespGetCheckInSchedulist{}

func (x RespGetCheckInSchedulist) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecGetCheckInSchedulist(req ReqGetCheckInSchedulist) (RespGetCheckInSchedulist, error) {
	var resp RespGetCheckInSchedulist
	err := c.executeWXApiPost("/cgi-bin/checkin/getcheckinschedulist", req, &resp, true)
	if err != nil {
		return RespGetCheckInSchedulist{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetCheckInSchedulist{}, bizErr
	}
	return resp, nil
}
