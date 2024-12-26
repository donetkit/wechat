package apis

import (
	"encoding/json"
)

type ReqGetCheckInOption struct {
	Datetime   int      `json:"datetime"`
	Useridlist []string `json:"useridlist"`
}

var _ bodyer = ReqGetCheckInOption{}

func (x ReqGetCheckInOption) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespGetCheckInOption struct {
	CommonResp
	Info []struct {
		Userid string `json:"userid"`
		Group  struct {
			Grouptype     int  `json:"grouptype"`
			Groupid       int  `json:"groupid"`
			OpenSpCheckin bool `json:"open_sp_checkin"`
			Checkindate   []struct {
				Workdays    []int `json:"workdays"`
				Checkintime []struct {
					WorkSec          int `json:"work_sec"`
					OffWorkSec       int `json:"off_work_sec"`
					RemindWorkSec    int `json:"remind_work_sec"`
					RemindOffWorkSec int `json:"remind_off_work_sec"`
				} `json:"checkintime"`
				FlexTime        int  `json:"flex_time"`
				NoneedOffwork   bool `json:"noneed_offwork"`
				LimitAheadtime  int  `json:"limit_aheadtime"`
				FlexOnDutyTime  int  `json:"flex_on_duty_time"`
				FlexOffDutyTime int  `json:"flex_off_duty_time"`
			} `json:"checkindate"`
			SpeWorkdays []struct {
				Timestamp   int    `json:"timestamp"`
				Notes       string `json:"notes"`
				Checkintime []struct {
					WorkSec          int `json:"work_sec"`
					OffWorkSec       int `json:"off_work_sec"`
					RemindWorkSec    int `json:"remind_work_sec"`
					RemindOffWorkSec int `json:"remind_off_work_sec"`
				} `json:"checkintime"`
			} `json:"spe_workdays"`
			SpeOffdays []struct {
				Timestamp   int           `json:"timestamp"`
				Notes       string        `json:"notes"`
				Checkintime []interface{} `json:"checkintime"`
			} `json:"spe_offdays"`
			SyncHolidays bool   `json:"sync_holidays"`
			Groupname    string `json:"groupname"`
			NeedPhoto    bool   `json:"need_photo"`
			WifimacInfos []struct {
				Wifiname string `json:"wifiname"`
				Wifimac  string `json:"wifimac"`
			} `json:"wifimac_infos"`
			NoteCanUseLocalPic     bool `json:"note_can_use_local_pic"`
			AllowCheckinOffworkday bool `json:"allow_checkin_offworkday"`
			AllowApplyOffworkday   bool `json:"allow_apply_offworkday"`
			LocInfos               []struct {
				Lat       int    `json:"lat"`
				Lng       int    `json:"lng"`
				LocTitle  string `json:"loc_title"`
				LocDetail string `json:"loc_detail"`
				Distance  int    `json:"distance"`
			} `json:"loc_infos"`
			Schedulelist []struct {
				ScheduleId   int    `json:"schedule_id"`
				ScheduleName string `json:"schedule_name"`
				TimeSection  []struct {
					TimeId           int  `json:"time_id"`
					WorkSec          int  `json:"work_sec"`
					OffWorkSec       int  `json:"off_work_sec"`
					RemindWorkSec    int  `json:"remind_work_sec"`
					RemindOffWorkSec int  `json:"remind_off_work_sec"`
					RestBeginTime    int  `json:"rest_begin_time"`
					RestEndTime      int  `json:"rest_end_time"`
					AllowRest        bool `json:"allow_rest"`
				} `json:"time_section"`
				LimitAheadtime  int  `json:"limit_aheadtime"`
				NoneedOffwork   bool `json:"noneed_offwork"`
				LimitOfftime    int  `json:"limit_offtime"`
				FlexOnDutyTime  int  `json:"flex_on_duty_time"`
				FlexOffDutyTime int  `json:"flex_off_duty_time"`
				AllowFlex       bool `json:"allow_flex"`
				LateRule        struct {
					AllowOffworkAfterTime bool `json:"allow_offwork_after_time"`
					Timerules             []struct {
						OffworkAfterTime int `json:"offwork_after_time"`
						OnworkFlexTime   int `json:"onwork_flex_time"`
					} `json:"timerules"`
				} `json:"late_rule"`
				MaxAllowArriveEarly int `json:"max_allow_arrive_early"`
				MaxAllowArriveLate  int `json:"max_allow_arrive_late"`
			} `json:"schedulelist"`
		} `json:"group"`
	} `json:"info"`
}

var _ bodyer = RespGetCheckInOption{}

func (x RespGetCheckInOption) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecGetCheckInOption(req ReqGetCheckInOption) (RespGetCheckInOption, error) {
	var resp RespGetCheckInOption
	err := c.executeWXApiPost("/cgi-bin/checkin/getcheckinoption", req, &resp, true)
	if err != nil {
		return RespGetCheckInOption{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetCheckInOption{}, bizErr
	}
	return resp, nil
}
