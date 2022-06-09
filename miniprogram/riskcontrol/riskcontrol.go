package riskcontrol

import (
	"context"
	"fmt"
	context2 "github.com/donetkit/wechat/miniprogram/context"
	"github.com/donetkit/wechat/util"
)

const (
	getUserRiskRankURL = "https://api.weixin.qq.com/wxa/getuserriskrank?access_token=%s"
)

// RiskControl 安全风控
type RiskControl struct {
	*context2.Context
}

// NewRiskControl init
func NewRiskControl(ctx *context2.Context) *RiskControl {
	return &RiskControl{ctx}
}

// UserRiskRankRequest 获取用户安全等级请求
type UserRiskRankRequest struct {
	AppID    string `json:"appid"`     // 小程序 app id
	OpenID   string `json:"openid"`    // 用户的 openid
	Scene    uint8  `json:"scene"`     // 场景值，0:注册，1:营销作弊
	ClientIP string `json:"client_ip"` // 用户访问源ip

	Mobile       string `json:"mobile_no"`     // 用户手机号
	Email        string `json:"email_address"` // 用户邮箱地址
	ExtendedInfo string `json:"extended_info"` // 额外补充信息
	IsTest       bool   `json:"is_test"`       // false：正式调用，true：测试调用
}

// UserRiskRank 用户安全等级
type UserRiskRank struct {
	util.CommonError
	UnionID  int64 `json:"union_id"`  // 唯一请求标识
	RiskRank uint8 `json:"risk_rank"` // 用户风险等级
}

// GetUserRiskRank 根据提交的用户信息数据获取用户的安全等级 risk_rank，无需用户授权。
func (riskControl *RiskControl) GetUserRiskRank(ctx context.Context, in *UserRiskRankRequest) (res UserRiskRank, err error) {
	accessToken, err := riskControl.GetAccessToken(ctx)
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getUserRiskRankURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &res, "GetUserRiskRank")
	return
}
