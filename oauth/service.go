package oauth

import "github.com/hiscaler/lingxing"

type service struct {
	lingXing *lingxing.LingXing
}

type Service interface {
	Auth(appId, appSecret string) (ar AuthResponse, err error)
	Refresh(appId, refreshToken string) (ar AuthResponse, err error)
}

func NewService(lx *lingxing.LingXing) Service {
	return service{lx}
}
