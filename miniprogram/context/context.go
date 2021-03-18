package context

import (
	"github.com/donetkit/wechat/credential"
	"github.com/donetkit/wechat/miniprogram/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
