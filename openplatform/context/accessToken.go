//Package context 开放平台相关context
package context

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/donetkit/wechat/util"
)

const (
	componentAccessTokenURL = "https://api.weixin.qq.com/cgi-bin/component/api_component_token"
	getPreCodeURL           = "https://api.weixin.qq.com/cgi-bin/component/api_create_preauthcode?component_access_token=%s"
	queryAuthURL            = "https://api.weixin.qq.com/cgi-bin/component/api_query_auth?component_access_token=%s"
	refreshTokenURL         = "https://api.weixin.qq.com/cgi-bin/component/api_authorizer_token?component_access_token=%s"
	getComponentInfoURL     = "https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_info?component_access_token=%s"
	componentLoginURL       = "https://mp.weixin.qq.com/cgi-bin/componentloginpage?component_appid=%s&pre_auth_code=%s&redirect_uri=%s"
	// componentLoginURL= "https://mp.weixin.qq.com/cgi-bin/componentloginpage?component_appid=%s&pre_auth_code=%s&redirect_uri=%s&auth_type=%d&biz_appid=%s"
	bindComponentURL = "https://mp.weixin.qq.com/safe/bindcomponent?action=bindcomponent&auth_type=%d&no_scan=1&component_appid=%s&pre_auth_code=%s&redirect_uri=%s&biz_appid=%s#wechat_redirect"
	//TODO 获取授权方选项信息
	//getComponentConfigURL = "https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_option?component_access_token=%s"
	//TODO 获取已授权的账号信息
	//getuthorizerListURL = "POST https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_list?component_access_token=%s"

	// 开放平台 AccessToken
	ComponentAccessTokenCacheKey = "WeiXin:Container:Open:AccessToken:%s"
	// 开放平台授权 公众号/小程序 AccessToken
	AuthorizerAccessTokenCacheKey = "WeiXin:Container:Open:AuthorizerAccessToken:%s"
)

// GetComponentAccessToken 获取 ComponentAccessToken
func (c *Context) GetComponentAccessToken(ctx context.Context) (string, error) {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	accessTokenCacheKey := fmt.Sprintf(ComponentAccessTokenCacheKey, c.AppID)
	val := c.Cache.WithContext(ctx).Get(accessTokenCacheKey)
	if val == nil {
		return "", fmt.Errorf("cann't get component access token")
	}
	return val.(string), nil
}

// SetComponentAccessToken 通过component_verify_ticket 获取 ComponentAccessToken
func (c *Context) SetComponentAccessToken(ctx context.Context, verifyTicket string) (*ComponentAccessToken, error) {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	body := map[string]string{
		"component_appid":         c.AppID,
		"component_appsecret":     c.AppSecret,
		"component_verify_ticket": verifyTicket,
	}
	respBody, err := util.PostJSON(componentAccessTokenURL, body)
	if err != nil {
		return nil, err
	}

	at := &ComponentAccessToken{}
	if err := json.Unmarshal(respBody, at); err != nil {
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
	body, err := util.PostJSON(uri, req)
	if err != nil {
		return "", err
	}

	var ret struct {
		PreCode string `json:"pre_auth_code"`
	}
	if err := json.Unmarshal(body, &ret); err != nil {
		return "", err
	}

	return ret.PreCode, nil
}

// GetComponentLoginPage 获取第三方公众号授权链接(扫码授权)
func (c *Context) GetComponentLoginPage(ctx context.Context, redirectURI string, authType int, bizAppID string) (string, error) {
	code, err := c.GetPreCode(ctx)
	if err != nil {
		return "", err
	}

	componentLoginURLNew := fmt.Sprintf(componentLoginURL, c.AppID, code, url.QueryEscape(redirectURI)) // , authType
	if bizAppID != "" {
		componentLoginURLNew = fmt.Sprintf("%s&biz_appid=%s", componentLoginURLNew, bizAppID)
	}
	return componentLoginURLNew, nil
}

// GetBindComponentURL 获取第三方公众号授权链接(链接跳转，适用移动端)
func (c *Context) GetBindComponentURL(ctx context.Context, redirectURI string, authType int, bizAppID string) (string, error) {
	code, err := c.GetPreCode(ctx)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(bindComponentURL, authType, c.AppID, code, url.QueryEscape(redirectURI), bizAppID), nil
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
	body, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, err
	}

	var ret struct {
		util.CommonError
		Info *AuthBaseInfo `json:"authorization_info"`
	}

	if err := json.Unmarshal(body, &ret); err != nil {
		return nil, err
	}
	if ret.ErrCode != 0 {
		err = fmt.Errorf("QueryAuthCode error : errcode=%v , errmsg=%v", ret.ErrCode, ret.ErrMsg)
		return nil, err
	}
	accessTokenCacheKey := fmt.Sprintf(AuthorizerAccessTokenCacheKey, ret.Info.Appid)
	authorizerAccessToken := &AuthorizerAccessToken{
		AuthorizationInfoExpireTime: time.Now().Unix() + ExpiryTimeSpan(ret.Info.ExpiresIn),
		AuthorizerAccessToken:       ret.Info.AuthrAccessToken,
	}
	val, _ := json.Marshal(authorizerAccessToken)
	if err := c.Cache.WithContext(ctx).Set(accessTokenCacheKey, string(val), 31*24*3600*time.Second); err != nil {
		return nil, nil
	}
	return ret.Info, nil
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
	body, err := util.PostJSON(uri, req)
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
	val := c.Cache.WithContext(ctx).Get(authTokenKey)
	if val == nil {
		return "", fmt.Errorf("cannot get authorizer %s access token", appId)
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
	if err := c.Cache.WithContext(ctx).Set(authTokenKey, string(res), 31*24*3600*time.Second); err != nil {
		return "", err
	}
	return ret.AccessToken, nil
}

// GetAuthAccessToken 获取授权方AccessToken
func (c *Context) GetAuthAccessToken(ctx context.Context, appId string) (string, error) {
	authorizerAccessToken := AuthorizerAccessToken{}
	authTokenKey := fmt.Sprintf(AuthorizerAccessTokenCacheKey, appId)
	val := c.Cache.WithContext(ctx).Get(authTokenKey)
	if val == nil {
		return "", fmt.Errorf("cannot get authorizer %s access token", appId)
	}
	if err := json.Unmarshal([]byte(val.(string)), &authorizerAccessToken); err != nil {
		return "", err
	}
	if authorizerAccessToken.AuthorizationInfoExpireTime < time.Now().Unix() {
		return c.refreshAuthToken(ctx, appId, authorizerAccessToken.AuthorizerAccessToken.RefreshToken)
	}
	return authorizerAccessToken.AuthorizerAccessToken.AccessToken, nil
}

// GetAuthInfo 获取授权方的帐号基本信息
func (c *Context) GetAuthInfo(ctx context.Context, appId string) (*AuthorizerInfo, *AuthBaseInfo, error) {
	cat, err := c.GetComponentAccessToken(ctx)
	if err != nil {
		return nil, nil, err
	}

	req := map[string]string{
		"component_appid":  c.AppID,
		"authorizer_appid": appId,
	}

	uri := fmt.Sprintf(getComponentInfoURL, cat)
	body, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, nil, err
	}

	var ret struct {
		AuthorizerInfo    *AuthorizerInfo `json:"authorizer_info"`
		AuthorizationInfo *AuthBaseInfo   `json:"authorization_info"`
	}
	if err := json.Unmarshal(body, &ret); err != nil {
		return nil, nil, err
	}

	return ret.AuthorizerInfo, ret.AuthorizationInfo, nil
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
