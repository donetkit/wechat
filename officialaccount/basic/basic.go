package basic

import (
	"context"
	"fmt"

	context2 "github.com/donetkit/wechat/officialaccount/context"
	"github.com/donetkit/wechat/util"
)

var (
	//获取微信服务器IP地址
	//文档：https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_the_WeChat_server_IP_address.html
	getCallbackIPURL  = "https://api.weixin.qq.com/cgi-bin/getcallbackip"
	getAPIDomainIPURL = "https://api.weixin.qq.com/cgi-bin/get_api_domain_ip"

	//清理接口调用次数
	clearQuotaURL = "https://api.weixin.qq.com/cgi-bin/clear_quota"
)

//Basic struct
type Basic struct {
	*context2.Context
}

//NewBasic 实例
func NewBasic(context *context2.Context) *Basic {
	basic := new(Basic)
	basic.Context = context
	return basic
}

//IPListRes 获取微信服务器IP地址 返回结果
type IPListRes struct {
	util.CommonError
	IPList []string `json:"ip_list"`
}

//GetCallbackIP 获取微信callback IP地址
func (basic *Basic) GetCallbackIP(ctx context.Context) ([]string, error) {
	ak, err := basic.GetAccessToken(ctx)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s?access_token=%s", getCallbackIPURL, ak)
	data, err := util.HTTPGet(url)
	if err != nil {
		return nil, err
	}
	ipListRes := &IPListRes{}
	err = util.DecodeWithError(data, ipListRes, "GetCallbackIP")
	return ipListRes.IPList, err
}

//GetAPIDomainIP 获取微信API接口 IP地址
func (basic *Basic) GetAPIDomainIP(ctx context.Context) ([]string, error) {
	ak, err := basic.GetAccessToken(ctx)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s?access_token=%s", getAPIDomainIPURL, ak)
	data, err := util.HTTPGet(url)
	if err != nil {
		return nil, err
	}
	ipListRes := &IPListRes{}
	err = util.DecodeWithError(data, ipListRes, "GetAPIDomainIP")
	return ipListRes.IPList, err
}

//ClearQuota 清理接口调用次数
func (basic *Basic) ClearQuota(ctx context.Context) error {
	ak, err := basic.GetAccessToken(ctx)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("%s?access_token=%s", clearQuotaURL, ak)
	data, err := util.PostJSON(url, map[string]string{
		"appid": basic.AppID,
	})
	if err != nil {
		return err
	}
	return util.DecodeWithCommonError(data, "ClearQuota")
}
