package apis

import (
	"net/url"
)

// 获取授权链接
var authUrlBaseCustomize = "https://open.weixin.qq.com/connect/oauth2/authorize"

type GetCustomizeOauthUrlReq struct {
	RedirectUri string `json:"redirect_uri"` // 授权回调跳转地址
	Scope       string `json:"scope"`        // 应用授权作用域
	State       string `json:"state"`        // 重定向后会带上state参数，企业可以填写a-zA-Z0-9的参数值，长度不可超过128个字节
	AgentId     string `json:"agent_id"`     // 应用agentid，建议填上该参数（对于第三方应用和代开发自建应用，在填写该参数的情况下或者在工作台、聊天工具栏、应用会话内发起oauth2请求的场景中，会触发接口许可的自动激活）。snsapi_privateinfo时必填否则报错；
}

// ?appid=CORPID&redirect_uri=REDIRECT_URI&response_type=code&scope=snsapi_base&state=STATE&agentid=AGENTID#wechat_redirect

// 构造oauth2链接
func (c *ApiClient) GetCustomizeOauthUrl(req GetCustomizeOauthUrlReq) string {
	ret, _ := url.Parse(authUrlBaseCustomize)
	query := ret.Query()
	query.Set("appid", c.CorpId)
	query.Set("redirect_uri", url.QueryEscape(req.RedirectUri))
	query.Set("response_type", "code")
	query.Set("scope", req.Scope)
	if req.State != "" {
		query.Set("state", req.State)
	}
	if req.AgentId != "" {
		query.Set("agentid", req.AgentId)
	}
	ret.RawQuery = query.Encode()
	return ret.String() + "#wechat_redirect"
}

//type GetJsSdkSignResp struct {
//	CorpId    string `json:"corp_id"`
//	Noncestr  string `json:"noncestr"`
//	Timestamp int64  `json:"timestamp"`
//	Signature string `json:"signature"`
//	AgentId   int64  `json:"agent_id,omitempty"`
//}
//
//// 获取前端 JS-SDK 使用权限签名
//// 文档：https://developer.work.weixin.qq.com/document/path/90506
//func (c *ApiClient) GetJsSdkSign(corpId, link, jsapiTicket string, agentId int) GetJsSdkSignResp {
//	var (
//		noncestr  = getRandomString(16)
//		timestamp = time.Now().Unix()
//	)
//
//	unescapeUrl, _ := url.QueryUnescape(link)
//
//	signature := newSha1(fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s", jsapiTicket, noncestr, timestamp, unescapeUrl))
//
//	return GetJsSdkSignResp{
//		CorpId:    corpId,
//		Noncestr:  noncestr,
//		Timestamp: timestamp,
//		Signature: signature,
//		AgentId:   int64(agentId),
//	}
//}
//
//func newSha1(s string) string {
//	h := sha1.New()
//	h.Write([]byte(s))
//	return fmt.Sprintf("%x", h.Sum(nil))
//}
//
//func getRandomString(n int) string {
//	rand.Seed(time.Now().UnixNano())
//	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
//	bytes := []byte(str)
//	var result []byte
//	for i := 0; i < n; i++ {
//		result = append(result, bytes[rand.Intn(len(bytes))])
//	}
//	return string(result)
//}
