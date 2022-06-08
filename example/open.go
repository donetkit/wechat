package main

import (
	"github.com/donetkit/contrib-log/glog"
	"github.com/donetkit/contrib/server/webserve"
	"github.com/gin-gonic/gin"
)

func main() {
	log := glog.New()
	gin.SetMode("debug")
	appServe := webserve.New(webserve.WithLogger(log))
	//routers.InitRouter(appServe)
	appServe.Run()

}
