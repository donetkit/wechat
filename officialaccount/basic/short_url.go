package basic

import (
	"context"
	"fmt"

	"github.com/donetkit/wechat/util"
)

const (
	// 将一条长链接转成短链接
	// https://developers.weixin.qq.com/doc/offiaccount/Account_Management/URL_Shortener.html
	long2shortURL    = "https://api.weixin.qq.com/cgi-bin/shorturl?access_token=%s"
	long2shortAction = "long2short"
)

type (
	reqLong2ShortURL struct {
		Action  string `json:"action"`
		LongURL string `json:"long_url"`
	}
	resplong2ShortURL struct {
		ShortURL string `json:"short_url"`
		util.CommonError
	}
)

// Long2ShortURL 将一条长链接转成短链接
func (basic *Basic) Long2ShortURL(ctx context.Context, longURL string) (shortURL string, err error) {
	var (
		req = &reqLong2ShortURL{
			Action:  long2shortAction,
			LongURL: longURL,
		}
		resp          = new(resplong2ShortURL)
		ac, uri       string
		responseBytes []byte
	)
	ac, err = basic.GetAccessTokenContext(ctx)
	if err != nil {
		return
	}
	uri = fmt.Sprintf(long2shortURL, ac)
	responseBytes, err = util.PostJSONContext(ctx, uri, req)
	if err != nil {
		return
	}
	err = util.DecodeWithError(responseBytes, resp, long2shortAction)
	return resp.ShortURL, err
}
