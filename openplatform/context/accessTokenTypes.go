package context

import "github.com/donetkit/wechat/util"

// ID 微信返回接口中各种类型字段
type ID struct {
	ID int `json:"id"`
}

// AuthBaseInfo 授权的基本信息
type AuthBaseInfo struct {
	AuthrAccessToken
	FuncInfo []AuthFuncInfo `json:"func_info"`
}

// AuthFuncInfo 授权的接口内容
type AuthFuncInfo struct {
	FuncscopeCategory ID `json:"funcscope_category"`
}

// AuthrAccessToken 授权方AccessToken
type AuthrAccessToken struct {
	Appid        string `json:"authorizer_appid"`
	AccessToken  string `json:"authorizer_access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"authorizer_refresh_token"`
}

// AuthorizerInfo 授权方详细信息
type AuthorizerInfo struct {
	NickName        string `json:"nick_name"`
	HeadImg         string `json:"head_img"`
	ServiceTypeInfo ID     `json:"service_type_info"`
	VerifyTypeInfo  ID     `json:"verify_type_info"`
	UserName        string `json:"user_name"`
	PrincipalName   string `json:"principal_name"`
	BusinessInfo    struct {
		OpenStore string `json:"open_store"`
		OpenScan  string `json:"open_scan"`
		OpenPay   string `json:"open_pay"`
		OpenCard  string `json:"open_card"`
		OpenShake string `json:"open_shake"`
	}
	Alias     string `json:"alias"`
	QrcodeURL string `json:"qrcode_url"`
}

// ComponentAccessToken 第三方平台
type ComponentAccessToken struct {
	util.CommonError
	AccessToken string `json:"component_access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

type AuthorizerAccessToken struct {
	AuthorizationInfoExpireTime int64            `json:"authorization_info_expire_time"`
	AuthorizerAccessToken       AuthrAccessToken `json:"authorizer_info"`
}
