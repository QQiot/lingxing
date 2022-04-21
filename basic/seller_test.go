package basic

import (
	"github.com/hiscaler/gox/jsonx"
	"testing"
)

func TestService_Sellers(t *testing.T) {
	items, err := lxService.Sellers()
	if err != nil {
		t.Errorf("lxService.Sellers error: %s", err.Error())
	} else {
		t.Log(jsonx.ToJson(items, "[]"))
	}
}
