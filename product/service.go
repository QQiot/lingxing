package product

import "github.com/hiscaler/lingxing"

type service struct {
	lingXing *lingxing.LingXing
}

type Service interface {
	Products(params ProductsQueryParams) (items []Product, nextOffset int, isLastPage bool, err error) // 本地产品列表
}

func NewService(lx *lingxing.LingXing) Service {
	return service{lx}
}
