package sale

import (
	"github.com/hiscaler/gox/jsonx"
	"testing"
)

func TestService_AmazonFBMOrders(t *testing.T) {
	params := AmazonFBMOrdersQueryParams{
		StartTime: "2022-01-01 00:00:00",
		EndTime:   "2022-01-01 23:59:59",
		SID:       "168",
	}
	items, _, _, err := lxService.AmazonFBMOrders(params)
	if err != nil {
		t.Errorf("lxService.AmazonFBMOrders error: %s", err.Error())
	} else {
		t.Log(jsonx.ToJson(items, "[]"))
	}
}
