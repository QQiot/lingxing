package basic

import "github.com/hiscaler/lingxing"

type service struct {
	lingXing *lingxing.LingXing
}

type Service interface {
	Sellers() (items []Seller, err error)                                     // 亚马逊店铺信息
	Accounts() (items []Account, err error)                                   // ERP账号列表
	Rates(params RatesQueryParams) (items []Rate, isLastPage bool, err error) // 汇率
}

func NewService(lx *lingxing.LingXing) Service {
	return service{lx}
}
