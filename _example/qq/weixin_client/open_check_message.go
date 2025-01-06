package weixin_client

import (
	"strings"

	"github.com/donetkit/wechat/officialaccount/message"
)

func OnTextMessage(msg message.MixMessage) *message.Reply {
	if msg.Content == "TESTCOMPONENT_MSG_TYPE_TEXT" {
		text := message.NewText(msg.Content + "_callback") //固定为TESTCOMPONENT_MSG_TYPE_TEXT_callback
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	}

	if strings.HasPrefix(msg.Content, "QUERY_AUTH_CODE:") {
		//authCode := strings.ReplaceAll(msg.Content, "QUERY_AUTH_CODE:", "")
		//authInfo, _ := OpenPlatformClient.QueryAuthCode(authCode)
		////authInfo.AccessToken
		//OpenPlatformClient.GetOfficialAccount(msg.AppID).GetCustomerMessageManager().Send()
		return nil
	}
	return nil
}

func OnEventMessage(msg message.MixMessage) *message.Reply {
	text := message.NewText(string(msg.Event) + "from_callback")
	return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
}
