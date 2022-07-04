package lingxing

import (
	"github.com/hiscaler/gox/jsonx"
	"testing"
)

func TestBasicDataService_Rates(t *testing.T) {
	params := RatesQueryParams{Date: "2021-01"}
	items, _, _, err := lingXingClient.Services.BasicData.Rates(params)
	if err != nil {
		t.Errorf("Services.BasicData.Rates() error: %s", err.Error())
	} else {
		t.Log(jsonx.ToJson(items, "[]"))
	}
}
