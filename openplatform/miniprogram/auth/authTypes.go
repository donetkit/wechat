package auth

// MiniProgramSessionKey 小程序
type MiniProgramSessionKey struct {
	OpenId     string `json:"open_id"`
	UnionId    string `json:"union_id"`
	SessionKey string `json:"session_key"`
	ExpireTime int64  `json:"expire_time"`
}
