package sale

import (
	"github.com/hiscaler/lingxing"
)

type service struct {
	lingXing *lingxing.LingXing
}

type Service interface {
	AmazonOrders(params AmazonOrdersQueryParams) (items []AmazonOrder, nextOffset int, isLastPage bool, err error)          // 亚马逊订单列表
	AmazonOrder(params AmazonOrderQueryParams) (detail AmazonOrderDetail, err error)                                        // 亚马逊订单详情
	AmazonFBMOrders(params AmazonFBMOrdersQueryParams) (items []AmazonFBMOrder, nextOffset int, isLastPage bool, err error) // 亚马逊自发货订单（FBM）列表
}

func NewService(lx *lingxing.LingXing) Service {
	return service{lx}
}
