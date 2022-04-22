package sale

import (
	"encoding/json"
	"fmt"
	"github.com/hiscaler/gox/jsonx"
	"github.com/hiscaler/lingxing"
	"github.com/hiscaler/lingxing/config"
	"os"
	"testing"
)

var lxInstance *lingxing.LingXing
var lxService Service

func TestMain(m *testing.M) {
	b, err := os.ReadFile("../config/config_test.json")
	if err != nil {
		panic(fmt.Sprintf("Read config error: %s", err.Error()))
	}
	var c config.Config
	err = json.Unmarshal(b, &c)
	if err != nil {
		panic(fmt.Sprintf("Parse config file error: %s", err.Error()))
	}
	lxInstance = lingxing.NewLingXing(c)
	lxService = NewService(lxInstance)
	m.Run()
}

func TestService_AmazonOrders(t *testing.T) {
	params := AmazonOrdersQueryParams{
		StartDate: "2022-01-01 00:00:00",
		EndDate:   "2022-01-01 23:59:59",
		SID:       168,
	}
	items, _, _, err := lxService.AmazonOrders(params)
	if err != nil {
		t.Errorf("lxService.AmazonOrders error: %s", err.Error())
	} else {
		t.Log(jsonx.ToJson(items, "[]"))
	}
}
