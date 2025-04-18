package auth

import (
	"context"
	"encoding/json"
	"fmt"
	context2 "github.com/donetkit/wechat/qqminiprogram/context"
	"github.com/donetkit/wechat/util"
)

const (
	code2SessionURL = "https://api.q.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"

	checkEncryptedDataURL = "https://api.q.qq.com/wxa/business/checkencryptedmsg?access_token=%s"
)

// Auth 登录/用户信息
type Auth struct {
	*context2.Context
}

// NewAuth new auth
func NewAuth(ctx *context2.Context) *Auth {
	return &Auth{ctx}
}

// ResCode2Session 登录凭证校验的返回结果
type ResCode2Session struct {
	util.CommonError

	OpenID     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` // 会话密钥
	UnionID    string `json:"unionid"`     // 用户在开放平台的唯一标识符，在满足UnionID下发条件的情况下会返回
}

// RspCheckEncryptedData .
type RspCheckEncryptedData struct {
	util.CommonError

	Vaild      bool `json:"vaild"`       // 是否是合法的数据
	CreateTime uint `json:"create_time"` // 加密数据生成的时间戳
}

// Code2Session 登录凭证校验。
func (auth *Auth) Code2Session(ctx context.Context, jsCode string) (result ResCode2Session, err error) {
	return auth.Code2SessionContext(ctx, jsCode)
}

// Code2SessionContext 登录凭证校验。
func (auth *Auth) Code2SessionContext(ctx context.Context, jsCode string) (result ResCode2Session, err error) {
	var response []byte
	if response, err = util.HTTPGetContext(ctx, fmt.Sprintf(code2SessionURL, auth.AppID, auth.AppSecret, jsCode)); err != nil {
		return
	}
	if err = json.Unmarshal(response, &result); err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("Code2Session error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}

// GetPaidUnionID 用户支付完成后，获取该用户的 UnionId，无需用户授权
func (auth *Auth) GetPaidUnionID() {
	// TODO
}

// CheckEncryptedData .检查加密信息是否由微信生成（当前只支持手机号加密数据），只能检测最近3天生成的加密数据
func (auth *Auth) CheckEncryptedData(ctx context.Context, encryptedMsgHash string) (result RspCheckEncryptedData, err error) {
	return auth.CheckEncryptedDataContext(ctx, encryptedMsgHash)
}

// CheckEncryptedDataContext .检查加密信息是否由微信生成（当前只支持手机号加密数据），只能检测最近3天生成的加密数据
func (auth *Auth) CheckEncryptedDataContext(ctx context.Context, encryptedMsgHash string) (result RspCheckEncryptedData, err error) {
	var response []byte
	var (
		at string
	)
	if at, err = auth.GetAccessTokenContext(ctx); err != nil {
		return
	}
	if response, err = util.HTTPPostContext(ctx, fmt.Sprintf(checkEncryptedDataURL, at), []byte("encrypted_msg_hash="+encryptedMsgHash), nil); err != nil {
		return
	}
	if err = util.DecodeWithError(response, &result, "CheckEncryptedDataAuth"); err != nil {
		return
	}
	return
}
