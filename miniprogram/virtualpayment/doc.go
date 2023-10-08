// Package virtualpayment mini program virtual payment
package virtualpayment

import (
	"github.com/donetkit/wechat/miniprogram/context"
)

// NewVirtualPayment 实例化小程序虚拟支付 API
func NewVirtualPayment(ctx *context.Context) *VirtualPayment {
	return &VirtualPayment{
		ctx: ctx,
	}
}
