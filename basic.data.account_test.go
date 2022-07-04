package lingxing

import (
	"github.com/hiscaler/gox/jsonx"
	"testing"
)

func TestBasicDataService_Accounts(t *testing.T) {
	items, err := lingXingClient.Services.BasicData.Accounts()
	if err != nil {
		t.Errorf("Services.BasicData.Accounts() error: %s", err.Error())
	} else {
		t.Log(jsonx.ToPrettyJson(items))
	}
}
