package apis

import (
	"encoding/json"
)

type ReqGetCheckInDayData struct {
	Starttime  int      `json:"starttime"`
	Endtime    int      `json:"endtime"`
	Useridlist []string `json:"useridlist"`
}

var _ bodyer = ReqGetCheckInDayData{}

func (x ReqGetCheckInDayData) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespGetCheckInDayData struct {
	CommonResp
	Datas []struct {
		BaseInfo struct {
			Date        int    `json:"date"`
			RecordType  int    `json:"record_type"`
			Name        string `json:"name"`
			NameEx      string `json:"name_ex"`
			DepartsName string `json:"departs_name"`
			Acctid      string `json:"acctid"`
			RuleInfo    struct {
				Groupid      int    `json:"groupid"`
				Groupname    string `json:"groupname"`
				Scheduleid   int    `json:"scheduleid"`
				Schedulename string `json:"schedulename"`
				Checkintime  []struct {
					WorkSec    int `json:"work_sec"`
					OffWorkSec int `json:"off_work_sec"`
				} `json:"checkintime"`
			} `json:"rule_info"`
			DayType int `json:"day_type"`
		} `json:"base_info"`
		SummaryInfo struct {
			CheckinCount    int `json:"checkin_count"`
			RegularWorkSec  int `json:"regular_work_sec"`
			StandardWorkSec int `json:"standard_work_sec"`
			EarliestTime    int `json:"earliest_time"`
			LastestTime     int `json:"lastest_time"`
		} `json:"summary_info"`
		HolidayInfos []struct {
			SpDescription struct {
				Data []struct {
					Lang string `json:"lang"`
					Text string `json:"text"`
				} `json:"data"`
			} `json:"sp_description"`
			SpNumber string `json:"sp_number"`
			SpTitle  struct {
				Data []struct {
					Lang string `json:"lang"`
					Text string `json:"text"`
				} `json:"data"`
			} `json:"sp_title"`
		} `json:"holiday_infos"`
		ExceptionInfos []struct {
			Count     int `json:"count"`
			Duration  int `json:"duration"`
			Exception int `json:"exception"`
		} `json:"exception_infos"`
		OtInfo struct {
			OtStatus           int   `json:"ot_status"`
			OtDuration         int   `json:"ot_duration"`
			ExceptionDuration  []int `json:"exception_duration"`
			WorkdayOverAsMoney int   `json:"workday_over_as_money"`
		} `json:"ot_info"`
		SpItems []struct {
			Count      int    `json:"count"`
			Duration   int    `json:"duration"`
			TimeType   int    `json:"time_type"`
			Type       int    `json:"type"`
			VacationId int    `json:"vacation_id"`
			Name       string `json:"name"`
		} `json:"sp_items"`
	} `json:"datas"`
}

var _ bodyer = RespGetCheckInDayData{}

func (x RespGetCheckInDayData) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecGetCheckInDayData(req ReqGetCheckInDayData) (RespGetCheckInDayData, error) {
	var resp RespGetCheckInDayData
	err := c.executeWXApiPost("/cgi-bin/checkin/getcheckin_daydata", req, &resp, true)
	if err != nil {
		return RespGetCheckInDayData{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetCheckInDayData{}, bizErr
	}
	return resp, nil
}
