package product

import "github.com/hiscaler/lingxing"

type service struct {
	lingXing *lingxing.LingXing
}

type Service interface {
	Products(params ProductsQueryParams) (items []Product, nextOffset int, isLastPage bool, err error) // 本地产品列表
	Product(id int) (item ProductDetail, err error)                                                    // 本地产品详情
	Brands(params BrandsQueryParams) (items []Brand, nextOffset int, isLastPage bool, err error)       // 本地产品品牌列表
}

func NewService(lx *lingxing.LingXing) Service {
	return service{lx}
}
