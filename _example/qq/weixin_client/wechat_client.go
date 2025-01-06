package weixin_client

import (
	"github.com/donetkit/contrib_cache/memory"
	"github.com/donetkit/wechat"
	"github.com/donetkit/wechat/openplatform"
	"qq/config"
	"sync"

	"github.com/donetkit/contrib-log/glog"

	openplatformConfig "github.com/donetkit/wechat/openplatform/config"
)

var OpenPlatformClient *openplatform.OpenPlatform

func InitWeiXin(log glog.ILogger) {
	wechatClient := wechat.NewWechat()

	//caChe := redis.New(redis.WithLogger(log), redis.WithAddr(config.ServerConfig.Redis.Addr), redis.WithPort(config.ServerConfig.Redis.Port), redis.WithPassword(config.ServerConfig.Redis.Password), redis.WithDB(config.ServerConfig.Redis.DB)).WithDB(config.ServerConfig.Redis.DB)

	caChe := memory.New().WithDB(2)

	cfg := &openplatformConfig.Config{
		AppID:          config.ServerConfig.OpenWechat.AppId,
		AppSecret:      config.ServerConfig.OpenWechat.AppSecret,
		Token:          config.ServerConfig.OpenWechat.Token,
		EncodingAESKey: config.ServerConfig.OpenWechat.EncodingAESKey,
		Cache:          caChe,
		Lock:           new(sync.Mutex),
	}
	OpenPlatformClient = wechatClient.GetOpenPlatform(cfg)
}
