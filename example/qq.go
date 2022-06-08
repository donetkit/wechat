package main

import (
	"github.com/donetkit/contrib-log/glog"
	"github.com/donetkit/contrib/db/redis"
	"github.com/donetkit/contrib/server/webserve"
	"github.com/donetkit/wechat"
	qqminiConfig "github.com/donetkit/wechat/qqminiprogram/config"
	"github.com/gin-gonic/gin"
)

func main() {
	log := glog.New()
	wc := wechat.NewWechat()
	cfg := &qqminiConfig.Config{
		AppID:     "xxx",
		AppSecret: "xxx",
		Cache:     redis.New(redis.WithLogger(log), redis.WithAddr("127.0.0.1"), redis.WithPort(6379), redis.WithDB(0)),
	}
	program := wc.GetQQMiniProgram(cfg)
	program.GetAuth()
	gin.SetMode("debug")
	appServe := webserve.New(webserve.WithLogger(log))
	//routers.InitRouter(appServe)
	appServe.Run()

}
