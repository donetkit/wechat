package apis

import (
	"encoding/json"
)

type ReqGetCheckInMonthData struct {
	Starttime  int      `json:"starttime"`
	Endtime    int      `json:"endtime"`
	Useridlist []string `json:"useridlist"`
}

var _ bodyer = ReqGetCheckInMonthData{}

func (x ReqGetCheckInMonthData) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespGetCheckInMonthData struct {
	CommonResp
	Datas []struct {
		BaseInfo struct {
			RecordType  int    `json:"record_type"`
			Name        string `json:"name"`
			NameEx      string `json:"name_ex"`
			DepartsName string `json:"departs_name"`
			RuleInfo    struct {
				Groupid   int    `json:"groupid"`
				Groupname string `json:"groupname"`
			} `json:"rule_info"`
			Acctid string `json:"acctid"`
		} `json:"base_info"`
		SummaryInfo struct {
			ExceptDays      int `json:"except_days"`
			RegularWorkSec  int `json:"regular_work_sec"`
			StandardWorkSec int `json:"standard_work_sec"`
			WorkDays        int `json:"work_days"`
		} `json:"summary_info"`
		ExceptionInfos []struct {
			Count     int `json:"count"`
			Duration  int `json:"duration"`
			Exception int `json:"exception"`
		} `json:"exception_infos"`
		SpItems []struct {
			Count      int    `json:"count"`
			Duration   int    `json:"duration"`
			TimeType   int    `json:"time_type"`
			Type       int    `json:"type"`
			VacationId int    `json:"vacation_id"`
			Name       string `json:"name"`
		} `json:"sp_items"`
		OverworkInfo struct {
			WorkdayOverSec         int `json:"workday_over_sec"`
			RestdaysOverSec        int `json:"restdays_over_sec"`
			WorkdaysOverAsVacation int `json:"workdays_over_as_vacation"`
			WorkdaysOverAsMoney    int `json:"workdays_over_as_money"`
			RestdaysOverAsVacation int `json:"restdays_over_as_vacation"`
			RestdaysOverAsMoney    int `json:"restdays_over_as_money"`
			HolidaysOverAsVacation int `json:"holidays_over_as_vacation"`
			HolidaysOverAsMoney    int `json:"holidays_over_as_money"`
		} `json:"overwork_info"`
	} `json:"datas"`
}

var _ bodyer = RespGetCheckInMonthData{}

func (x RespGetCheckInMonthData) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecGetCheckInMonthData(req ReqGetCheckInMonthData) (RespGetCheckInMonthData, error) {
	var resp RespGetCheckInMonthData
	err := c.executeWXApiPost("/cgi-bin/checkin/getcheckin_monthdata", req, &resp, true)
	if err != nil {
		return RespGetCheckInMonthData{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetCheckInMonthData{}, bizErr
	}
	return resp, nil
}
