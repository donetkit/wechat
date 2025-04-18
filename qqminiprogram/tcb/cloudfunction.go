package tcb

import (
	"context"
	"fmt"

	"github.com/donetkit/wechat/util"
)

const (
	//触发云函数
	invokeCloudFunctionURL = "https://api.q.qq.com/tcb/invokecloudfunction"
)

// InvokeCloudFunctionRes 云函数调用返回结果
type InvokeCloudFunctionRes struct {
	util.CommonError
	RespData string `json:"resp_data"` //云函数返回的buffer
}

// InvokeCloudFunction 云函数调用
//
//reference:https://developers.weixin.qq.com/miniprogram/dev/wxcloud/reference-http-api/functions/invokeCloudFunction.html
func (tcb *Tcb) InvokeCloudFunction(ctx context.Context, env, name, args string) (*InvokeCloudFunctionRes, error) {
	accessToken, err := tcb.GetAccessTokenContext(ctx)
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("%s?access_token=%s&env=%s&name=%s", invokeCloudFunctionURL, accessToken, env, name)
	response, err := util.HTTPPost(uri, args)
	if err != nil {
		return nil, err
	}
	invokeCloudFunctionRes := &InvokeCloudFunctionRes{}
	err = util.DecodeWithError(response, invokeCloudFunctionRes, "InvokeCloudFunction")
	return invokeCloudFunctionRes, err
}
