package credential

import "context"

// JsTicketHandle js ticket获取
type JsTicketHandle interface {
	//GetTicket 获取ticket
	GetTicket(ctx context.Context, accessToken string) (ticket string, err error)
}
