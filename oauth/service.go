package oauth

import "github.com/hiscaler/lingxing"

type service struct {
	lingXing *lingxing.LingXing
}

type Service interface {
	Auth() (ar AuthResponse, err error)
	Refresh() (ar AuthResponse, err error)
}

func NewService(lx *lingxing.LingXing) Service {
	return service{lx}
}
