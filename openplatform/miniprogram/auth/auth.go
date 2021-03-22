package auth

import (
	"encoding/json"
	"fmt"
	"time"

	openContext "github.com/donetkit/wechat/openplatform/context"
	"github.com/donetkit/wechat/openplatform/miniprogram/encryptor"

	"github.com/donetkit/wechat/util"
)

const (
	code2SessionURL = "https://api.weixin.qq.com/sns/component/jscode2session?appid=%s&js_code=%s&grant_type=authorization_code&component_appid=%s&component_access_token=%s"

	// 开放平台授权 小程序 SessionKey
	MiniProgramSessionKeyCacheKey = "WeiXin:Container:Open:MiniProgramSessionKey:%s:%s"
)

//Auth 登录/用户信息
type Auth struct {
	*openContext.Context
	appID string // 小程序appId
}

//NewAuth new auth
func NewAuth(opContext *openContext.Context, appID string) *Auth {
	return &Auth{Context: opContext, appID: appID}
}

// ResCode2Session 登录凭证校验的返回结果
type ResCode2Session struct {
	util.CommonError
	OpenID     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` // 会话密钥
	UnionID    string `json:"unionid"`     // 用户在开放平台的唯一标识符，在满足UnionID下发条件的情况下会返回
}

//Code2Session 登录凭证校验。
func (auth *Auth) Code2Session(jsCode string) (result ResCode2Session, err error) {
	componentAK, err := auth.GetComponentAccessToken()
	urlStr := fmt.Sprintf(code2SessionURL, auth.appID, jsCode, auth.Context.AppID, componentAK)
	var response []byte
	response, err = util.HTTPGet(urlStr)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("Code2Session error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	miniProgramSessionKey := MiniProgramSessionKey{
		OpenId:     result.OpenID,
		UnionId:    result.UnionID,
		SessionKey: result.SessionKey,
	}
	res, err := json.Marshal(miniProgramSessionKey)
	if err != nil {
		err = fmt.Errorf("json Marshal %s SessionKey", "")
		return
	}
	miniProgramSessionKeyCacheKey := fmt.Sprintf(MiniProgramSessionKeyCacheKey, auth.appID, result.OpenID)
	auth.Cache.Set(miniProgramSessionKeyCacheKey, string(res), 15*24*3600*time.Second)
	return
}

//GetSessionKey 登录SessionKey。
func (auth *Auth) GetSessionKey(openID string) (result string, err error) {
	miniProgramSessionKeyCacheKey := fmt.Sprintf(MiniProgramSessionKeyCacheKey, auth.appID, openID)
	val := auth.Cache.Get(miniProgramSessionKeyCacheKey)
	if val == nil {
		return "", fmt.Errorf("SessionKey Cache %s", "")
	}
	miniProgramSessionKey := MiniProgramSessionKey{}
	err = json.Unmarshal([]byte(val.(string)), &miniProgramSessionKey)
	if err != nil {
		return "", fmt.Errorf("json Marshal Session Key %s", "")
	}
	return miniProgramSessionKey.SessionKey, nil
}

func (auth *Auth) GetDecryptData(sessionKey, encryptedData, iv string) (*encryptor.PlainData, error) {
	return encryptor.NewEncryptor(auth.Context, auth.appID).Decrypt(sessionKey, encryptedData, iv)
}

//GetPaidUnionID 用户支付完成后，获取该用户的 UnionId，无需用户授权
func (auth *Auth) GetPaidUnionID() {
	//TODO
}
