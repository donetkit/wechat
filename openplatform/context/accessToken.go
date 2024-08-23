package context

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/donetkit/wechat/util"
	"net/url"
	"time"
)

const (
	componentAccessTokenURL = "https://api.weixin.qq.com/cgi-bin/component/api_component_token"
	getPreCodeURL           = "https://api.weixin.qq.com/cgi-bin/component/api_create_preauthcode?component_access_token=%s"
	queryAuthURL            = "https://api.weixin.qq.com/cgi-bin/component/api_query_auth?component_access_token=%s"
	refreshTokenURL         = "https://api.weixin.qq.com/cgi-bin/component/api_authorizer_token?component_access_token=%s"
	getComponentInfoURL     = "https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_info?component_access_token=%s"
	componentLoginURL       = "https://mp.weixin.qq.com/cgi-bin/componentloginpage?component_appid=%s&pre_auth_code=%s&redirect_uri=%s&auth_type=%d&biz_appid=%s"
	bindComponentURL        = "https://mp.weixin.qq.com/safe/bindcomponent?action=bindcomponent&auth_type=%d&no_scan=1&component_appid=%s&pre_auth_code=%s&redirect_uri=%s&biz_appid=%s#wechat_redirect"
	// TODO 获取授权方选项信息
	// getComponentConfigURL = "https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_option?component_access_token=%s"
	// TODO 获取已授权的账号信息
	// getuthorizerListURL = "POST https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_list?component_access_token=%s"

	// 开放平台 AccessToken
	ComponentAccessTokenCacheKey = "WeiXin:Container:Open:AccessToken:%s"
	// 开放平台授权 公众号/小程序 AccessToken
	AuthorizerAccessTokenCacheKey = "WeiXin:Container:Open:AuthorizerAccessToken:%s"
)

// ComponentAccessToken 第三方平台
type ComponentAccessToken struct {
	util.CommonError
	AccessToken string `json:"component_access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

// GetComponentAccessToken 获取 ComponentAccessToken
func (c *Context) GetComponentAccessToken(ctx context.Context) (string, error) {
	accessTokenCacheKey := fmt.Sprintf(ComponentAccessTokenCacheKey, c.AppID)
	val := c.Cache.WithContext(ctx).Get(accessTokenCacheKey)
	if val == nil {
		return "", fmt.Errorf("cann't get component access token")
	}
	return val.(string), nil
}

// SetComponentAccessToken 通过component_verify_ticket 获取 ComponentAccessToken
func (c *Context) SetComponentAccessToken(ctx context.Context, verifyTicket string) (*ComponentAccessToken, error) {
	body := map[string]string{
		"component_appid":         c.AppID,
		"component_appsecret":     c.AppSecret,
		"component_verify_ticket": verifyTicket,
	}
	respBody, err := util.PostJSONContext(ctx, componentAccessTokenURL, body)
	if err != nil {
		return nil, err
	}

	at := &ComponentAccessToken{}
	if err := util.DecodeWithError(respBody, at, "SetComponentAccessToken"); err != nil {
		return nil, err
	}

	if at.ErrCode != 0 {
		return nil, fmt.Errorf("SetComponentAccessToken Error , errcode=%d , errmsg=%s", at.ErrCode, at.ErrMsg)
	}

	accessTokenCacheKey := fmt.Sprintf(ComponentAccessTokenCacheKey, c.AppID)
	expires := at.ExpiresIn - 1500
	if err := c.Cache.WithContext(ctx).Set(accessTokenCacheKey, at.AccessToken, time.Duration(expires)*time.Second); err != nil {
		return nil, nil
	}
	return at, nil
}

// GetPreCode 获取预授权码
func (c *Context) GetPreCode(ctx context.Context) (string, error) {
	cat, err := c.GetComponentAccessToken(ctx)
	if err != nil {
		return "", err
	}
	req := map[string]string{
		"component_appid": c.AppID,
	}
	uri := fmt.Sprintf(getPreCodeURL, cat)
	body, err := util.PostJSONContext(ctx, uri, req)
	if err != nil {
		return "", err
	}

	var ret struct {
		PreCode string `json:"pre_auth_code"`
	}
	err = json.Unmarshal(body, &ret)
	return ret.PreCode, err
}

// GetComponentLoginPage 获取第三方公众号授权链接(扫码授权)
func (c *Context) GetComponentLoginPage(ctx context.Context, redirectURI string, authType int, bizAppID string) (string, error) {
	code, err := c.GetPreCode(ctx)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(componentLoginURL, c.AppID, code, url.QueryEscape(redirectURI), authType, bizAppID), nil
}

// GetBindComponentURL 获取第三方公众号授权链接(链接跳转，适用移动端)
func (c *Context) GetBindComponentURL(ctx context.Context, redirectURI string, authType int, bizAppID string) (string, error) {
	code, err := c.GetPreCode(ctx)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(bindComponentURL, authType, c.AppID, code, url.QueryEscape(redirectURI), bizAppID), nil
}

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

// QueryAuthCode 使用授权码换取公众号或小程序的接口调用凭据和授权信息
func (c *Context) QueryAuthCode(ctx context.Context, authCode string) (*AuthBaseInfo, error) {
	cat, err := c.GetComponentAccessToken(ctx)
	if err != nil {
		return nil, err
	}

	req := map[string]string{
		"component_appid":    c.AppID,
		"authorization_code": authCode,
	}
	uri := fmt.Sprintf(queryAuthURL, cat)
	body, err := util.PostJSONContext(ctx, uri, req)
	if err != nil {
		return nil, err
	}

	var ret struct {
		util.CommonError
		Info *AuthBaseInfo `json:"authorization_info"`
	}

	if err := util.DecodeWithError(body, &ret, "QueryAuthCode"); err != nil {
		return nil, err
	}
	if ret.ErrCode != 0 {
		err = fmt.Errorf("QueryAuthCode error : errcode=%v , errmsg=%v", ret.ErrCode, ret.ErrMsg)
		return nil, err
	}
	return ret.Info, nil
}

// RefreshAuthrToken 获取（刷新）授权公众号或小程序的接口调用凭据（令牌）
func (c *Context) RefreshAuthrToken(ctx context.Context, appid, refreshToken string) (*AuthrAccessToken, error) {
	cat, err := c.GetComponentAccessToken(ctx)
	if err != nil {
		return nil, err
	}

	req := map[string]string{
		"component_appid":          c.AppID,
		"authorizer_appid":         appid,
		"authorizer_refresh_token": refreshToken,
	}
	uri := fmt.Sprintf(refreshTokenURL, cat)
	body, err := util.PostJSONContext(ctx, uri, req)
	if err != nil {
		return nil, err
	}

	ret := &AuthrAccessToken{}
	if err := util.DecodeWithError(body, ret, "RefreshAuthrToken"); err != nil {
		return nil, err
	}

	authTokenKey := fmt.Sprintf(AuthorizerAccessTokenCacheKey, appid)
	if err := c.Cache.WithContext(ctx).Set(authTokenKey, ret.AccessToken, time.Second*time.Duration(ret.ExpiresIn-30)); err != nil {
		return nil, err
	}
	return ret, nil
}

// GetAuthAccessToken 获取授权方AccessToken
func (c *Context) GetAuthAccessToken(ctx context.Context, appid string) (string, error) {
	authTokenKey := fmt.Sprintf(AuthorizerAccessTokenCacheKey, appid)
	val := c.Cache.WithContext(ctx).Get(authTokenKey)
	if val == nil {
		return "", fmt.Errorf("cannot get authorizer %s access token", appid)
	}
	return val.(string), nil
}

// GetAuthAccessToken1 获取授权方AccessToken
func (c *Context) GetAuthAccessToken1(ctx context.Context, appId string) (string, error) {
	authorizerAccessToken := AuthorizerAccessToken{}
	authTokenKey := fmt.Sprintf(AuthorizerAccessTokenCacheKey, appId)

	val := c.Cache.Get(authTokenKey)
	if val == nil {
		str, _ := c.Cache.GetString(authTokenKey)
		if len(str) > 0 {
			var reply interface{}
			json.Unmarshal([]byte(str), &reply)
			val = reply
		}
	}

	_, ok := val.(string)
	if !ok {
		v, err := json.Marshal(val)
		if err == nil {
			val = string(v)
		}
	}

	if err := json.Unmarshal([]byte(val.(string)), &authorizerAccessToken); err != nil {
		return "", fmt.Errorf("cannot get authorizer %s access token 2", appId)
	}

	if authorizerAccessToken.AuthorizationInfoExpireTime < time.Now().Unix() {
		return c.refreshAuthToken(ctx, appId, authorizerAccessToken.AuthorizerAccessToken.RefreshToken)
	}
	return authorizerAccessToken.AuthorizerAccessToken.AccessToken, nil
}

// RefreshAuthToken 获取（刷新）授权公众号或小程序的接口调用凭据（令牌）
func (c *Context) refreshAuthToken(ctx context.Context, appId, refreshToken string) (string, error) {
	cat, err := c.GetComponentAccessToken(ctx)
	if err != nil {
		return "", err
	}

	req := map[string]string{
		"component_appid":          c.AppID,
		"authorizer_appid":         appId,
		"authorizer_refresh_token": refreshToken,
	}
	uri := fmt.Sprintf(refreshTokenURL, cat)
	body, err := util.PostJSONContext(ctx, uri, req)
	if err != nil {
		return "", err
	}

	ret := AuthrAccessToken{}
	if err := json.Unmarshal(body, &ret); err != nil {
		return "", err
	}
	ret.Appid = appId
	authTokenKey := fmt.Sprintf(AuthorizerAccessTokenCacheKey, appId)
	authorizerAccessToken := &AuthorizerAccessToken{}

	val := c.Cache.Get(authTokenKey)
	if val == nil {
		return "", fmt.Errorf("cannot get authorizer %s access token", appId)
	}

	_, ok := val.(string)
	if !ok {
		v, err := json.Marshal(val)
		if err == nil {
			val = string(v)
		}
	}

	if err := json.Unmarshal([]byte(val.(string)), &authorizerAccessToken); err != nil {
		return "", err
	}

	authorizerAccessToken.AuthorizerAccessToken = ret
	authorizerAccessToken.AuthorizationInfoExpireTime = time.Now().Unix() + ExpiryTimeSpan(ret.ExpiresIn)

	res, err := json.Marshal(authorizerAccessToken)
	if err != nil {
		return "", fmt.Errorf("json Marshal authorizer %s access token", appId)
	}
	if err := c.Cache.Set(authTokenKey, string(res), 5*12*31*24*3600*time.Second); err != nil {
		return "", err
	}
	return ret.AccessToken, nil
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

	MiniProgramInfo *MiniProgramInfo       `json:"MiniProgramInfo"`
	RegisterType    int                    `json:"register_type"`
	AccountStatus   int                    `json:"account_status"`
	BasicConfig     *AuthorizerBasicConfig `json:"basic_config"`
}

// AuthorizerBasicConfig 授权账号的基础配置结构体
type AuthorizerBasicConfig struct {
	IsPhoneConfigured bool `json:"isPhoneConfigured"`
	IsEmailConfigured bool `json:"isEmailConfigured"`
}

// MiniProgramInfo 授权账号小程序配置 授权账号为小程序时存在
type MiniProgramInfo struct {
	Network struct {
		RequestDomain   []string `json:"RequestDomain"`
		WsRequestDomain []string `json:"WsRequestDomain"`
		UploadDomain    []string `json:"UploadDomain"`
		DownloadDomain  []string `json:"DownloadDomain"`
		BizDomain       []string `json:"BizDomain"`
		UDPDomain       []string `json:"UDPDomain"`
	} `json:"network"`
	Categories []CategoriesInfo `json:"categories"`
}

// CategoriesInfo 授权账号小程序配置的类目信息
type CategoriesInfo struct {
	First  string `wx:"first"`
	Second string `wx:"second"`
}

// GetAuthrInfo 获取授权方的帐号基本信息
func (c *Context) GetAuthrInfo(ctx context.Context, appid string) (*AuthorizerInfo, *AuthBaseInfo, error) {
	cat, err := c.GetComponentAccessToken(ctx)
	if err != nil {
		return nil, nil, err
	}

	req := map[string]string{
		"component_appid":  c.AppID,
		"authorizer_appid": appid,
	}

	uri := fmt.Sprintf(getComponentInfoURL, cat)
	body, err := util.PostJSONContext(ctx, uri, req)
	if err != nil {
		return nil, nil, err
	}

	var ret struct {
		AuthorizerInfo    *AuthorizerInfo `json:"authorizer_info"`
		AuthorizationInfo *AuthBaseInfo   `json:"authorization_info"`
	}
	if err := util.DecodeWithError(body, &ret, "GetAuthrInfo"); err != nil {
		return nil, nil, err
	}

	return ret.AuthorizerInfo, ret.AuthorizationInfo, nil
}

type AuthorizerAccessToken struct {
	AuthorizationInfoExpireTime int64            `json:"authorization_info_expire_time"`
	AuthorizerAccessToken       AuthrAccessToken `json:"authorizer_info"`
}

func ExpiryTimeSpan(expireInSeconds int64) int64 {
	if expireInSeconds > 3600 {
		expireInSeconds -= 600
	} else if expireInSeconds > 1800 {
		expireInSeconds -= 300
	} else if expireInSeconds > 1800 {
		expireInSeconds -= 30
	}
	return expireInSeconds
}
