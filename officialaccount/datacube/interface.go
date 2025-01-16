package datacube

import (
	"context"
	"fmt"

	"github.com/donetkit/wechat/util"
)

const (
	getInterfaceSummary     = "https://api.weixin.qq.com/datacube/getinterfacesummary"
	getInterfaceSummaryHour = "https://api.weixin.qq.com/datacube/getinterfacesummaryhour"
)

// ResInterfaceSummary 接口分析数据响应
type ResInterfaceSummary struct {
	util.CommonError

	List []struct {
		RefDate       string `json:"ref_date"`
		CallbackCount int    `json:"callback_count"`
		FailCount     int    `json:"fail_count"`
		TotalTimeCost int    `json:"total_time_cost"`
		MaxTimeCost   int    `json:"max_time_cost"`
	} `json:"list"`
}

// ResInterfaceSummaryHour 接口分析分时数据响应
type ResInterfaceSummaryHour struct {
	util.CommonError

	List []struct {
		RefDate       string `json:"ref_date"`
		RefHour       int    `json:"ref_hour"`
		CallbackCount int    `json:"callback_count"`
		FailCount     int    `json:"fail_count"`
		TotalTimeCost int    `json:"total_time_cost"`
		MaxTimeCost   int    `json:"max_time_cost"`
	} `json:"list"`
}

// GetInterfaceSummary 获取接口分析数据
func (cube *DataCube) GetInterfaceSummary(ctx context.Context, s string, e string) (resInterfaceSummary ResInterfaceSummary, err error) {
	accessToken, err := cube.GetAccessTokenContext(ctx)
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", getInterfaceSummary, accessToken)
	reqDate := &reqDate{
		BeginDate: s,
		EndDate:   e,
	}

	response, err := util.PostJSONContext(ctx, uri, reqDate)
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &resInterfaceSummary, "GetInterfaceSummary")
	return
}

// GetInterfaceSummaryHour 获取接口分析分时数据
func (cube *DataCube) GetInterfaceSummaryHour(ctx context.Context, s string, e string) (resInterfaceSummaryHour ResInterfaceSummaryHour, err error) {
	accessToken, err := cube.GetAccessTokenContext(ctx)
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", getInterfaceSummaryHour, accessToken)
	reqDate := &reqDate{
		BeginDate: s,
		EndDate:   e,
	}

	response, err := util.PostJSONContext(ctx, uri, reqDate)
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &resInterfaceSummaryHour, "GetInterfaceSummaryHour")
	return
}
