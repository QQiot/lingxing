package basic

import (
	"github.com/hiscaler/gox/jsonx"
	"testing"
)

func TestService_Rates(t *testing.T) {
	params := RatesQueryParams{Date: "2021-01"}
	items, _, _, err := lxService.Rates(params)
	if err != nil {
		t.Errorf("lxService.Rates error: %s", err.Error())
	} else {
		t.Log(jsonx.ToJson(items, "[]"))
	}
}
