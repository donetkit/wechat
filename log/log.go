package log

import "github.com/donetkit/contrib-log/glog"

var Log glog.ILoggerEntry

func InitLogger(logger glog.ILogger) {
	Log = logger.WithField("Open-WeiXin", "Open-WeiXin")
}
