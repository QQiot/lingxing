package sale

import (
	"github.com/hiscaler/lingxing"
)

type service struct {
	lingXing *lingxing.LingXing
}

type Service interface {
	AmazonOrders(params AmazonOrdersQueryParams) (items []AmazonOrder, nextOffset int, isLastPage bool, err error) // 亚马逊订单列表
}

func NewService(lx *lingxing.LingXing) Service {
	return service{lx}
}
