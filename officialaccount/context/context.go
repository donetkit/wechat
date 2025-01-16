package context

import (
	"github.com/donetkit/wechat/credential"
	"github.com/donetkit/wechat/officialaccount/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenContextHandle
}
