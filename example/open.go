package main

import (
	"github.com/donetkit/contrib-log/glog"
	"github.com/donetkit/contrib/server/webserve"
	"github.com/gin-gonic/gin"
)

func main() {
	log := glog.New()
	log.Info("======================================================================")
	gin.SetMode("debug")
	logs := glog.New()
	appServe := webserve.New(webserve.WithLogger(logs))
	//routers.InitRouter(appServe)
	appServe.Run()

}
