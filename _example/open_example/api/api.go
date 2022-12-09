package api

import (
	"encoding/json"
	"fmt"
	"github.com/donetkit/wechat/log"
	"github.com/donetkit/wechat/officialaccount/message"
	"net/http"
	"open_example/config"
	"open_example/weixin_client"

	"github.com/gin-gonic/gin"
)

const (
	OpenCallNoticeRedisKey = "WeiXinContainerOpenNotice"

	OpenCallBackRedisKey = "WeiXinContainerOpenCallBack"
)

func OpenOauthJumpMp(c *gin.Context) {
	c.HTML(http.StatusOK, "jump_to_mp_oauth.tmpl", gin.H{
		"title": "跳转到公众号授权页",
	})
}

// OpenCallNotice 授权事件接收URL
func OpenCallNotice(c *gin.Context) {
	c.Writer.WriteString("success")
	server := weixin_client.OpenPlatformClient.GetServer(c, "")
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg *message.MixMessage) *message.Reply {
		if msg.InfoType != message.InfoTypeVerifyTicket {
			jsonData, err := json.Marshal(msg)
			if err == nil {
				server.Cache.WithContext(c.Request.Context()).XAdd(OpenCallNoticeRedisKey, "", []string{"", string(jsonData)})
			}
		}
		switch msg.InfoType {
		case message.InfoTypeVerifyTicket:
			weixin_client.OnComponentVerifyTicketRequest(c.Request.Context(), *msg)
		case message.InfoTypeAuthorized:
			weixin_client.OnAuthorizedRequest(*msg)
		case message.InfoTypeUnauthorized:
			weixin_client.OnUnauthorizedRequest(*msg)
		case message.InfoTypeUpdateAuthorized:
			weixin_client.OnUpdateAuthorizedRequest(*msg)
		default:
			log.Log.Info("未知的InfoType请求类型")
		}
		return nil
	})

	err := server.Serve() //处理消息接收以及回复
	if err != nil {
		log.Log.Error(err)
	}
}

// OpenCallBack 消息与事件接收URL
func OpenCallBack(c *gin.Context) {
	checkPublish := false
	appId := c.Param("appId")
	server := weixin_client.OpenPlatformClient.GetServer(c, appId)
	server.SetMessageHandler(func(msg *message.MixMessage) *message.Reply {
		jsonData, err := json.Marshal(msg)
		if err == nil {
			server.Cache.WithContext(c.Request.Context()).XAdd(OpenCallBackRedisKey, "", []string{appId, string(jsonData)})
		}
		if checkPublish {
			if msg.MsgType == message.MsgTypeText {
				return weixin_client.OnTextMessage(*msg)
			}
			if msg.MsgType == message.MsgTypeEvent {
				return weixin_client.OnEventMessage(*msg)
			}
		}
		return nil
	})
	err := server.Serve() //处理消息接收以及回复
	if err != nil {
		log.Log.Error(err)
	}
	server.Send()
}

// OpenOauth 发起授权页的体验URL
func OpenOauth(c *gin.Context) {
	var callbackUrl = fmt.Sprintf("%s/open/oauth/call/back", config.ServerConfig.OpenWechat.Url) //成功回调地址
	location, err := weixin_client.OpenPlatformClient.GetComponentLoginPage(c.Request.Context(), callbackUrl, 0, "")
	if err != nil {
		c.Writer.WriteString("fail")
		return
	}
	c.Redirect(http.StatusFound, location)
}

// OpenOauthCallBack 公众号授权成功回调地址  string auth_code, int expires_in, string appId
func OpenOauthCallBack(c *gin.Context) {
	authCode := c.Query("auth_code")
	if authCode != "" {
		_, err := weixin_client.OpenPlatformClient.QueryAuthCode(c.Request.Context(), authCode)
		if err != nil {
			c.Writer.WriteString(err.Error())
			return
		}
	}
	c.Writer.WriteString("success")

}
