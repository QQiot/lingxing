package lingxing

import (
	"github.com/hiscaler/gox/jsonx"
	"testing"
)

func TestBasicDataService_Sellers(t *testing.T) {
	items, err := lingXingClient.Services.BasicData.Sellers()
	if err != nil {
		t.Errorf("Services.BasicData.Sellers() error: %s", err.Error())
	} else {
		t.Log(jsonx.ToPrettyJson(items))
	}
}
