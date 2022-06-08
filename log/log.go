package log

import "github.com/donetkit/contrib-log/glog"

var Log glog.ILogger

func InitLogger(logger glog.ILogger) {
	Log = logger
}
