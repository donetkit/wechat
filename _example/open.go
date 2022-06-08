package main

import (
	"github.com/donetkit/contrib-log/glog"
	"github.com/donetkit/contrib/server/webserve"
	"github.com/donetkit/wechat/log"
	"github.com/gin-gonic/gin"
)

type name struct {
	name string
	age  int
}

func main() {
	logs := glog.New()
	log.InitLogger(logs)

	replyMsg := name{}
	log.Log.Debugf("response msg =%+v", replyMsg)

	gin.SetMode("debug")
	appServe := webserve.New(webserve.WithLogger(logs))
	//routers.InitRouter(appServe)
	appServe.Run()

}
