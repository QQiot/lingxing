package lingxing

import "time"

type Token struct {
	AccessToken     string `json:"access_token"`
	RefreshToken    string `json:"refresh_token"`
	ExpiresIn       int    `json:"expires_in"`
	ExpiresDatetime int64  `json:"expires_datetime"`
}

func (t Token) Valid() bool {
	return t.AccessToken != "" && t.RefreshToken != "" && t.ExpiresDatetime > time.Now().Unix()
}
