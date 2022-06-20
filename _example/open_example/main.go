package main

import (
	"github.com/donetkit/contrib-log/glog"
	"github.com/donetkit/contrib/server/webserve"
	"open_example/config"
	"open_example/router"
)

func main() {
	log := glog.New(glog.WithLevel(glog.Level(config.ServerConfig.App.LoggerLevel)))
	appServe := webserve.New(
		webserve.WithPort(config.ServerConfig.App.Port),
		webserve.WithLogger(log),
		webserve.WithProtocol("Restfull API"),
		webserve.WithReadTimeout(config.ServerConfig.App.ReadTimeOut),
		webserve.WithWriterTimeout(config.ServerConfig.App.WriterTimeOut))
	router.InitRouter(appServe, log)
	appServe.Run()
}
