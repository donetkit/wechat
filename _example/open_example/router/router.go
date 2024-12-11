package router

import (
	"github.com/donetkit/contrib-log/glog"
	"github.com/donetkit/contrib_gin_middleware/cors"
	"github.com/donetkit/contrib_gin_middleware/favicon"
	"github.com/donetkit/contrib_gin_middleware/logger"
	"github.com/donetkit/contrib_gin_middleware/requestid"
	"github.com/gin-gonic/gin"
	"net/http"
	"open_example/api"
	"open_example/weixin_client"

	"github.com/donetkit/contrib/server/webserve"
)

var regexEndpoints = []string{".*(/ts/)", ".*(/favicon.ico)", ".*(/health/)", "/metrics"}

func InitRouter(appServe *webserve.Server, log glog.ILogger) {
	weixin_client.InitWeiXin(log)
	logger.SetGinDefaultWriter(log)
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	if appServe.IsDevelopment() {
		r.Use(logger.New(logger.WithLogger(log), logger.WithExcludeRegexEndpoint(regexEndpoints)))
	}
	r.Use(cors.New(cors.WithExposeHeaders([]string{"X-Requested-With", "Accept"})))
	r.Use(favicon.New(favicon.WithRoutePaths("/favicon.ico", "./favicon.ico")))
	r.Use(requestid.New())
	r.LoadHTMLGlob("templates/*")
	r.Use(logger.NewErrorLogger(logger.WithLogger(log), logger.WithWriterErrorFn(func(c *gin.Context, logParams *logger.LogFormatterParams) (int, interface{}) {
		log.WithField("GIN-Exception", "GIN-Exception").Error(log)
		return http.StatusInternalServerError, "网络超时, 请重试!"
	})))

	if appServe.IsDevelopment() {
		gin.SetMode(gin.DebugMode)
	}
	r.GET("/", func(c *gin.Context) { c.String(200, "") })
	r.GET("/:path", func(c *gin.Context) {
		path := c.Param("path")
		c.File(path)
	})

	r.GET("/open/oauth/jump", api.OpenOauthJumpMp)        // 跳转到公众号授权页
	r.POST("/open/call/notice", api.OpenCallNotice)       // 授权事件接收URL
	r.POST("/open/call/back/:appId", api.OpenCallBack)    // 消息与事件接收URL
	r.GET("/open/oauth", api.OpenOauth)                   // 公众号授权跳转
	r.GET("/open/oauth/call/back", api.OpenOauthCallBack) // 公众号授权成功回调地址

	appServe.AddHandler(r)
}
