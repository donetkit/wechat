package weixin_client

import (
	"context"
	"log"

	"github.com/donetkit/wechat/officialaccount/message"
)

func OnComponentVerifyTicketRequest(ctx context.Context, msg message.MixMessage) *message.Reply {
	_, err := OpenPlatformClient.SetComponentAccessToken(ctx, msg.ComponentVerifyTicket)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func OnUnauthorizedRequest(msg message.MixMessage) *message.Reply {
	return nil
}

func OnAuthorizedRequest(msg message.MixMessage) *message.Reply {
	return nil
}

func OnUpdateAuthorizedRequest(msg message.MixMessage) *message.Reply {
	return nil
}

func OnThirdFasteRegisterRequest(msg message.MixMessage) *message.Reply {
	return nil
}

func OnNicknameAuditRequest(msg message.MixMessage) *message.Reply {
	return nil
}
