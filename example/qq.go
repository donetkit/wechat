package main

import (
	"fmt"
	"github.com/donetkit/wechat"
	"github.com/donetkit/wechat/cache"
	qqminiConfig "github.com/donetkit/wechat/qqminiprogram/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {

	wc := wechat.NewWechat()
	memory := cache.NewRedis(nil)
	cfg := &qqminiConfig.Config{
		AppID:     "xxx",
		AppSecret: "xxx",
		Cache:     memory,
	}
	program := wc.GetQQMiniProgram(cfg)

	program.GetAuth()

	fmt.Println("======================================================================")
	fmt.Println("main")
	fmt.Println("======================================================================")
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
