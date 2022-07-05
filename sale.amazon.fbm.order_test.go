package lingxing

import (
	"github.com/hiscaler/gox/jsonx"
	"testing"
)

func TestSaleService_AmazonFBMOrders(t *testing.T) {
	params := AmazonFBMOrdersQueryParams{
		StartTime: "2022-01-01 00:00:00",
		EndTime:   "2022-11-01 23:59:59",
		SID:       "172",
	}
	items, _, _, err := lingXingClient.Services.Sale.FBM.Order.All(params)
	if err != nil {
		t.Errorf("Services.Sale.FBM.Order.All() error: %s", err.Error())
	} else {
		t.Log(jsonx.ToPrettyJson(items))
	}
}
