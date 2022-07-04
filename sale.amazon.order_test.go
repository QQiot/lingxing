package lingxing

import (
	"github.com/hiscaler/gox/jsonx"
	"testing"
)

func TestSaleService_AmazonOrders(t *testing.T) {
	params := AmazonOrdersQueryParams{
		StartDate: "2022-01-01 00:00:00",
		EndDate:   "2022-01-01 23:59:59",
		SID:       168,
	}
	items, _, _, err := lingXingClient.Services.Sale.AmazonOrders(params)
	if err != nil {
		t.Errorf("Services.Sale.AmazonOrders() error: %s", err.Error())
	} else {
		t.Log(jsonx.ToPrettyJson(items))
	}
}

func TestSaleService_AmazonOrder(t *testing.T) {
	params := AmazonOrderQueryParams{
		OrderId: "123",
	}
	detail, err := lingXingClient.Services.Sale.AmazonOrder(params)
	if err != nil {
		t.Errorf("Services.Sale.AmazonOrder() error: %s", err.Error())
	} else {
		t.Log(jsonx.ToPrettyJson(detail))
	}
}