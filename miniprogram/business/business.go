package business

import "github.com/donetkit/wechat/miniprogram/context"

// Business 业务
type Business struct {
	*context.Context
}

// NewBusiness init
func NewBusiness(ctx *context.Context) *Business {
	return &Business{ctx}
}
