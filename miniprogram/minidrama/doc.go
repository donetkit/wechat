// Package minidrama Mini Program entertainment mini-drama related interface
package minidrama

import (
	"github.com/donetkit/wechat/miniprogram/context"
)

// NewMiniDrama 实例化小程序娱乐直播 API
func NewMiniDrama(ctx *context.Context) *MiniDrama {
	return &MiniDrama{
		ctx: ctx,
	}
}
