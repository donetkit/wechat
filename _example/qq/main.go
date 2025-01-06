package main

import (
	"github.com/donetkit/contrib-log/glog"
	"github.com/donetkit/contrib/server/webserve"
	"github.com/donetkit/contrib_cache/redis"
	"github.com/donetkit/wechat"
	"github.com/donetkit/wechat/log"
	qqminiConfig "github.com/donetkit/wechat/qqminiprogram/config"
	"github.com/gin-gonic/gin"
	"qq/config"
	"qq/router"
)

func main() {
	logs := glog.New()
	log.InitLogger(logs)
	wc := wechat.NewWechat()
	cfg := &qqminiConfig.Config{
		AppID:     "xxx",
		AppSecret: "xxx",
		Cache:     redis.New(redis.WithLogger(logs), redis.WithAddr("127.0.0.1"), redis.WithPort(6379), redis.WithDB(0)),
	}
	program := wc.GetQQMiniProgram(cfg)
	program.GetAuth()
	gin.SetMode("debug")
	appServe := webserve.New(
		webserve.WithPort(config.ServerConfig.App.Port),
		webserve.WithLogger(logs),
		webserve.WithProtocol("Restfull API"),
		webserve.WithReadTimeout(config.ServerConfig.App.ReadTimeOut),
		webserve.WithWriterTimeout(config.ServerConfig.App.WriterTimeOut))
	router.InitRouter(appServe, logs)
	appServe.Run()

}
