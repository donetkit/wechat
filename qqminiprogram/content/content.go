package content

import (
	"context"
	"fmt"
	context2 "github.com/donetkit/wechat/qqminiprogram/context"
	"github.com/donetkit/wechat/util"
)

const (
	checkTextURL  = "https://api.q.qq.com/wxa/msg_sec_check?access_token=%s"
	checkImageURL = "https://api.q.qq.com/wxa/img_sec_check?access_token=%s"
)

// Content 内容安全
type Content struct {
	*context2.Context
}

// NewContent 内容安全接口
func NewContent(ctx *context2.Context) *Content {
	return &Content{ctx}
}

// CheckText 检测文字
// @text 需要检测的文字
func (content *Content) CheckText(ctx context.Context, text string) error {
	accessToken, err := content.GetAccessTokenContext(ctx)
	if err != nil {
		return err
	}
	response, err := util.PostJSONContext(ctx,
		fmt.Sprintf(checkTextURL, accessToken),
		map[string]string{
			"content": text,
		},
	)
	if err != nil {
		return err
	}
	return util.DecodeWithCommonError(response, "ContentCheckText")
}

// CheckImage 检测图片
// 所传参数为要检测的图片文件的绝对路径，图片格式支持PNG、JPEG、JPG、GIF, 像素不超过 750 x 1334，同时文件大小以不超过 300K 为宜，否则可能报错
// @media 图片文件的绝对路径
func (content *Content) CheckImage(ctx context.Context, media string) error {
	accessToken, err := content.GetAccessTokenContext(ctx)
	if err != nil {
		return err
	}
	response, err := util.PostFile(
		"media",
		media,
		fmt.Sprintf(checkImageURL, accessToken),
	)
	if err != nil {
		return err
	}
	return util.DecodeWithCommonError(response, "ContentCheckImage")
}
