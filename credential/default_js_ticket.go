package credential

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/donetkit/contrib_cache/cache"
	"sync"
	"time"

	"github.com/donetkit/wechat/util"
)

// 获取ticket的url
const getTicketURL = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi"

// DefaultJsTicket 默认获取js ticket方法
type DefaultJsTicket struct {
	appID          string
	cacheKeyPrefix string
	cache          cache.ICache
	//jsAPITicket 读写锁 同一个AppID一个
	jsAPITicketLock *sync.Mutex
}

// NewDefaultJsTicket new
func NewDefaultJsTicket(appID string, cacheKeyPrefix string, cache cache.ICache) JsTicketHandle {
	return &DefaultJsTicket{
		appID:           appID,
		cache:           cache,
		cacheKeyPrefix:  cacheKeyPrefix,
		jsAPITicketLock: new(sync.Mutex),
	}
}

// ResTicket 请求jsapi_tikcet返回结果
type ResTicket struct {
	util.CommonError

	Ticket    string `json:"ticket"`
	ExpiresIn int64  `json:"expires_in"`
}

// GetTicket 获取jsapi_ticket
func (js *DefaultJsTicket) GetTicket(ctx context.Context, accessToken string) (ticketStr string, err error) {
	//先从cache中取
	jsAPITicketCacheKey := fmt.Sprintf("%s_jsapi_ticket_%s", js.cacheKeyPrefix, js.appID)

	ticketStr, err = js.cache.WithContext(ctx).GetString(jsAPITicketCacheKey)
	if len(ticketStr) > 0 {
		return
	}

	js.jsAPITicketLock.Lock()
	defer js.jsAPITicketLock.Unlock()

	// 双检，防止重复从微信服务器获取

	ticketStr, err = js.cache.WithContext(ctx).GetString(jsAPITicketCacheKey)
	if len(ticketStr) > 0 {
		return
	}

	var ticket ResTicket
	ticket, err = GetTicketFromServer(accessToken)
	if err != nil {
		return
	}
	expires := ticket.ExpiresIn - 1500
	err = js.cache.WithContext(ctx).Set(jsAPITicketCacheKey, ticket.Ticket, time.Duration(expires)*time.Second)
	ticketStr = ticket.Ticket
	return
}

// GetTicketFromServer 从服务器中获取ticket
func GetTicketFromServer(accessToken string) (ticket ResTicket, err error) {
	var response []byte
	url := fmt.Sprintf(getTicketURL, accessToken)
	response, err = util.HTTPGet(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &ticket)
	if err != nil {
		return
	}
	if ticket.ErrCode != 0 {
		err = fmt.Errorf("getTicket Error : errcode=%d , errmsg=%s", ticket.ErrCode, ticket.ErrMsg)
		return
	}
	return
}
