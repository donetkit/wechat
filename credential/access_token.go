package credential

import "context"

//AccessTokenHandle AccessToken 接口
type AccessTokenHandle interface {
	GetAccessToken(ctx context.Context) (accessToken string, err error)
}
