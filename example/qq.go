package main

import (
	"fmt"
	"github.com/donetkit/contrib-log/glog"
	"github.com/donetkit/contrib/db/redis"
	"github.com/donetkit/wechat"
	qqminiConfig "github.com/donetkit/wechat/qqminiprogram/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	log := glog.New()
	log.Info("======================================================================")
	log.Info("main")
	log.Info("======================================================================")
	wc := wechat.NewWechat()
	cacheClient := redis.New(redis.WithLogger(log), redis.WithAddr("127.0.0.1"), redis.WithPort(6379), redis.WithDB(0))
	cfg := &qqminiConfig.Config{
		AppID:     "xxx",
		AppSecret: "xxx",
		Cache:     cacheClient,
	}
	program := wc.GetQQMiniProgram(cfg)

	program.GetAuth()
	gin.SetMode("debug")
	//routersInit := routers.InitRouter()
	readTimeout := 30 * time.Second
	writerTimeout := 30 * time.Second
	endPoint := fmt.Sprintf(":%d", 80)
	maxHeaderBytes := 1 << 20
	server := &http.Server{
		Addr: endPoint,
		//Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writerTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	err := server.ListenAndServe()
	if err != nil {
		println(err.Error())
	}

}
