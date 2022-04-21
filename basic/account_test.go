package basic

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


func TestService_Sellers(t *testing.T) {
	sellers, err := lxService.Sellers()
	if err != nil {
		t.Errorf("lxService.Sellers error: %s", err.Error())
	} else {
		t.Log(jsonx.ToJson(sellers, "[]"))
	}
}
