package basic

import "github.com/hiscaler/lingxing"

type service struct {
	lingXing *lingxing.LingXing
}

type Service interface {
	Sellers() (items []Seller, err error)
}

func NewService(lx *lingxing.LingXing) Service {
	return service{lx}
}
