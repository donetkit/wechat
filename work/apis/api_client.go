package apis

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// DcsAppSuiteTicket 分布式app_suite_ticket：获取和设置suite_ticket的值，自行实现该接口的具体逻辑，比如使用redis方案【企微服务器每十分钟推送一次suite_ticket】
type DcsAppSuiteTicket interface {
	Get(cacheKey string) string                          // 获取suite_ticket
	Set(cacheKey, suiteTicket string, ttl time.Duration) // 设置suite_ticket
}

// ApiClient 企业微信客户端
type ApiClient struct {
	CorpId             string // 企业ID
	CorpProviderSecret string // 企业密钥

	// 第三方应用/代开发必填字段
	AppSuiteId     string // 应用唯一身份标识
	AppSuiteSecret string // 应用密钥
	AppSuiteTicket string // 企业微信服务器会定时（每十分钟）推送ticket。ticket会实时变更，并用于后续接口的调用。

	// 授权企业必填字段
	CompanyPermanentCode string // 企业授权给应用的永久授权码
	AgentId              int    // 授权方应用id

	accessTokenName        string // token参数名,默认access_token; 第三方应用为suite_access_token；第三方应用服务商为provider_access_token
	accessToken            *token
	jsapiTicket            *token
	jsapiTicketAgentConfig *token

	dcsSuiteTicketCacheKey string            // suite_ticket 缓存key，企微每十分钟更新一次
	dcsAppSuiteTicket      DcsAppSuiteTicket // 分布式app_suite_ticket

	ThirdAppClient *ApiClient // 第三方应用client，用于授权企业API客户端获取suite_access_token，目前用于第三方应用获取企业凭证接口

	logger Logger
}

// NewProviderApiClient 第三方服务商API客户端初始化
func NewProviderApiClient(corpId, corpProviderSecret string, opts Options) *ApiClient {
	c := ApiClient{
		CorpId:             corpId,
		CorpProviderSecret: corpProviderSecret,
		accessTokenName:    AccessTokenName,
		accessToken: &token{
			mutex:         &sync.RWMutex{},
			dcsToken:      opts.DcsToken,
			tokenCacheKey: fmt.Sprintf("%s#%s#%s", Provider, AccessTokenName, corpId),
		},
		dcsSuiteTicketCacheKey: fmt.Sprintf("%s#%s#%s", Provider, SuiteTicket, corpId),
		dcsAppSuiteTicket:      opts.DcsAppSuiteTicket,
		logger:                 opts.Logger,
	}

	if c.logger == nil {
		c.logger = loggerPrint{}
	}

	c.accessToken.setGetTokenFunc(c.getProviderAccessToken)

	return &c
}

// NewThirdAppApiClient 第三方应用API客户端初始化，第一次调用这个接口时，appSuiteTicket为空字符串
func NewThirdAppApiClient(corpId, appSuiteId, appSuiteSecret, appSuiteTicket string, opts Options) *ApiClient {
	c := ApiClient{
		//CorpId:          corpId,
		AppSuiteId:      appSuiteId,
		AppSuiteSecret:  appSuiteSecret,
		AppSuiteTicket:  appSuiteTicket,
		accessTokenName: SuiteAccessTokenName,
		accessToken: &token{
			mutex:         &sync.RWMutex{},
			dcsToken:      opts.DcsToken,
			tokenCacheKey: fmt.Sprintf("%s#%s#%s", ThirdApp, SuiteAccessTokenName, appSuiteId),
		},
		dcsSuiteTicketCacheKey: fmt.Sprintf("%s#%s#%s", ThirdApp, SuiteTicket, appSuiteId),
		dcsAppSuiteTicket:      opts.DcsAppSuiteTicket,
		logger:                 opts.Logger,
	}

	if c.logger == nil {
		c.logger = loggerPrint{}
	}

	c.accessToken.setGetTokenFunc(c.getSuiteToken)

	return &c
}

// NewCustomizedApiClient 自建应用代开发API客户端初始化，第一次调用这个接口时，appSuiteTicket为空字符串
func NewCustomizedApiClient(corpId, appSuiteId, appSuiteSecret, appSuiteTicket string, companyPermanentCode string, opts Options) *ApiClient {
	c := ApiClient{
		CorpId:          corpId,
		AppSuiteId:      appSuiteId,
		AppSuiteSecret:  appSuiteSecret,
		AppSuiteTicket:  appSuiteTicket,
		accessTokenName: AccessTokenName,
		accessToken: &token{
			mutex:         &sync.RWMutex{},
			dcsToken:      opts.DcsToken,
			tokenCacheKey: fmt.Sprintf("%s#%s#%s", CustomizedApp, SuiteAccessTokenName, appSuiteId),
		},
		dcsSuiteTicketCacheKey: fmt.Sprintf("%s#%s#%s", CustomizedApp, CustomizedTicket, appSuiteId),
		dcsAppSuiteTicket:      opts.DcsAppSuiteTicket,
		logger:                 opts.Logger,
		CompanyPermanentCode:   companyPermanentCode,
	}

	if c.logger == nil {
		c.logger = loggerPrint{}
	}

	c.accessToken.setGetTokenFunc(c.getCustomizedAuthCorpToken)

	return &c
}

// NewAuthCorpApiClient 第三方应用授权企业API客户端初始化
func NewAuthCorpApiClient(corpId, companyPermanentCode string, AgentId int, thirdAppClient *ApiClient, opts Options) *ApiClient {
	c := ApiClient{
		//CorpId:               corpId,
		AgentId:              AgentId,
		CompanyPermanentCode: companyPermanentCode,
		accessTokenName:      AccessTokenName,
		accessToken: &token{
			mutex:         &sync.RWMutex{},
			dcsToken:      opts.DcsToken,
			tokenCacheKey: fmt.Sprintf("%s#%s#%s", ThirdApp, AccessTokenName, thirdAppClient.AppSuiteId),
		},
		jsapiTicket: &token{
			mutex:         &sync.RWMutex{},
			dcsToken:      opts.DcsToken,
			tokenCacheKey: fmt.Sprintf("%s#%s#%s", ThirdApp, JsApiTicket, thirdAppClient.AppSuiteId),
		},
		jsapiTicketAgentConfig: &token{
			mutex:         &sync.RWMutex{},
			dcsToken:      opts.DcsToken,
			tokenCacheKey: fmt.Sprintf("%s#%s#%s", ThirdApp, JsApiTicketAgentConfig, thirdAppClient.AppSuiteId),
		},
		ThirdAppClient: thirdAppClient,
		logger:         opts.Logger,
	}

	if c.logger == nil {
		c.logger = loggerPrint{}
	}

	c.accessToken.setGetTokenFunc(c.getAuthCorpToken)

	c.jsapiTicket.setGetTokenFunc(c.getJSAPITicket)

	c.jsapiTicketAgentConfig.setGetTokenFunc(c.getJSAPITicketAgentConfig)

	return &c
}

// NewCustomizedAuthCorpApiClient 自建应用代开发授权企业API客户端初始化
func NewCustomizedAuthCorpApiClient(corpId, companyPermanentCode string, AgentId int, customizedAppClient *ApiClient, opts Options) *ApiClient {
	c := ApiClient{
		//CorpId:               corpId,
		AgentId:              AgentId,
		CompanyPermanentCode: companyPermanentCode,
		accessTokenName:      AccessTokenName,
		accessToken: &token{
			mutex:         &sync.RWMutex{},
			dcsToken:      opts.DcsToken,
			tokenCacheKey: fmt.Sprintf("%s#%s#%s", CustomizedApp, AccessTokenName, customizedAppClient.AppSuiteId),
		},
		logger: opts.Logger,
	}

	if c.logger == nil {
		c.logger = loggerPrint{}
	}

	c.accessToken.setGetTokenFunc(c.getCustomizedAuthCorpToken)

	return &c
}

func (c *ApiClient) composeWXApiURL(path string, req interface{}) *url.URL {
	values := url.Values{}
	if valuer, ok := req.(urlValuer); ok {
		values = valuer.intoURLValues()
	}

	base, err := url.Parse(DefaultQYAPIHost)
	if err != nil {
		panic(fmt.Sprintf("qyapiHost invalid: host=%s err=%+v", DefaultQYAPIHost, err))
	}

	base.Path = path
	base.RawQuery = values.Encode()

	return base
}

func (c *ApiClient) composeWXURLWithToken(path string, req interface{}, withAccessToken bool) *url.URL {
	wxApiURL := c.composeWXApiURL(path, req)

	if !withAccessToken {
		return wxApiURL
	}

	q := wxApiURL.Query()

	if len(c.CompanyPermanentCode) > 0 {
		q.Set(c.accessTokenName, c.accessToken.getToken())
	} else {
		if wxApiURL.Path == PathGetCorpToken { // 兼容获取企业凭证接口
			q.Set(SuiteAccessTokenName, c.ThirdAppClient.accessToken.getToken())
		} else {
			q.Set(c.accessTokenName, c.accessToken.getToken())
		}
	}

	//if wxApiURL.Path == PathGetCorpToken { // 兼容获取企业凭证接口
	//	q.Set(SuiteAccessTokenName, c.ThirdAppClient.accessToken.getToken())
	//} else {
	//	q.Set(c.accessTokenName, c.accessToken.getToken())
	//}

	wxApiURL.RawQuery = q.Encode()

	return wxApiURL
}

func (c *ApiClient) executeWXApiGet(path string, req urlValuer, objResp interface{}, withAccessToken bool) error {
	wxUrlWithToken := c.composeWXURLWithToken(path, req, withAccessToken)
	urlStr := wxUrlWithToken.String()

	httpReq := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(httpReq)

	httpResp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(httpResp)

	httpReq.SetRequestURI(urlStr)
	httpReq.Header.SetMethod(http.MethodGet)

	c.logger.Error("executeWXApiGet 请求Url:" + urlStr)
	c.logger.Errorf("executeWXApiGet 请求Url httpReq:%v", httpReq)
	if err := FastClient.DoTimeout(httpReq, httpResp, HttpTTL); err != nil {
		return err
	}
	respBody := httpResp.Body()
	c.logger.Errorf("executeWXApiGet 请求Url httpResp:%s", string(respBody))
	if len(respBody) == 0 { // 避免 json.Unmarshal 出现 unexpected end of JSON input 错误
		c.logger.Errorf("请求企微路由=%s, resp=%s, err=返回空响应体，建议降低并发数试试", urlStr, string(respBody))
		return errors.New("http resp body is nil")
	}
	go func() {
		defer func() {
			if err := recover(); err != nil {
				c.logger.Errorf("请求企微路由=%s, resp=%s, err=%+v", path, string(respBody), err)
			}
		}()
		c.RemoveTokenByHttpClient(respBody)
	}()

	return json.Unmarshal(respBody, &objResp)
}

func (c *ApiClient) executeWXApiPost(path string, req bodyer, objResp interface{}, withAccessToken bool) error {
	wxUrlWithToken := c.composeWXURLWithToken(path, req, withAccessToken)
	urlStr := wxUrlWithToken.String()

	reqBody, err := req.intoBody()
	if err != nil {
		return err
	}

	httpReq := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(httpReq)

	httpResp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(httpResp)

	httpReq.SetRequestURI(urlStr)
	httpReq.Header.SetContentType("application/json")
	httpReq.SetBody(reqBody)
	httpReq.Header.SetMethod(http.MethodPost)

	c.logger.Error("executeWXApiPost 请求Url:" + urlStr)
	c.logger.Errorf("executeWXApiPost 请求Url httpReq:%v", httpReq)
	if err := FastClient.DoTimeout(httpReq, httpResp, HttpTTL); err != nil {
		return err
	}

	respBody := httpResp.Body()
	c.logger.Errorf("executeWXApiPost 请求Url httpResp:%s", string(reqBody))
	if len(respBody) == 0 { // 避免 json.Unmarshal 出现 unexpected end of JSON input 错误
		c.logger.Errorf("请求企微路由=%s, req=%s, resp=%s, err=返回空响应体，建议降低并发数试试", path, string(reqBody), respBody)
		return errors.New("http resp body is nil")
	}

	go func() {
		defer func() {
			if err := recover(); err != nil {
				c.logger.Errorf("请求企微路由=%s, resp=%s, err=%+v", path, string(respBody), err)
			}
		}()
		c.RemoveTokenByHttpClient(respBody)
	}()

	return json.Unmarshal(respBody, &objResp)
}

func (c *ApiClient) executeWXApiMediaUpload(path string, req mediaUploader, objResp interface{}, withAccessToken bool) error {
	wxUrlWithToken := c.composeWXURLWithToken(path, req, withAccessToken)

	urlStr := wxUrlWithToken.String()

	m := req.getMedia()

	httpReq := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(httpReq)

	httpResp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(httpResp)

	// 新建一个缓冲，用于存放文件内容
	bodyBuffer := &bytes.Buffer{}
	// 创建一个multipart文件写入器，方便按照http规定格式写入内容
	bodyWriter := multipart.NewWriter(bodyBuffer)
	// 从bodyWriter生成fileWriter,并将文件内容写入fileWriter,多个文件可进行多次
	fileWriter, err := bodyWriter.CreateFormFile("media", m.filename)
	if err != nil {
		c.logger.Error(err.Error())
		return err
	}

	_, err = io.Copy(fileWriter, m.stream)
	if err != nil {
		return err
	}

	// 停止写入
	_ = bodyWriter.Close()

	httpReq.SetRequestURI(urlStr)
	httpReq.Header.SetContentType(bodyWriter.FormDataContentType())
	httpReq.SetBody(bodyBuffer.Bytes())
	httpReq.Header.SetMethod(http.MethodPost)

	if err := FastClient.DoTimeout(httpReq, httpResp, HttpTTL); err != nil {
		return err
	}

	respBody := httpResp.Body()
	if len(respBody) == 0 { // 避免 json.Unmarshal 出现 unexpected end of JSON input 错误
		c.logger.Errorf("请求企微路由=%s, resp=%s, err=返回空响应体，建议降低并发数试试", urlStr, string(respBody))
		return errors.New("http resp body is nil")
	}

	go func() {
		defer func() {
			if err := recover(); err != nil {
				c.logger.Errorf("请求企微路由=%s, resp=%s, err=%+v", urlStr, string(respBody), err)
			}
		}()
		c.RemoveTokenByHttpClient(respBody)
	}()

	return json.Unmarshal(respBody, &objResp)
}
