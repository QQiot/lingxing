package product

import (
	"fmt"
	"github.com/hiscaler/gox/jsonx"
	"github.com/hiscaler/lingxing"
	"github.com/hiscaler/lingxing/config"
	jsoniter "github.com/json-iterator/go"
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
	err = jsoniter.Unmarshal(b, &c)
	if err != nil {
		panic(fmt.Sprintf("Parse config file error: %s", err.Error()))
	}
	lxInstance = lingxing.NewLingXing(c)
	lxService = NewService(lxInstance)
	m.Run()
}

func TestService_Products(t *testing.T) {
	params := ProductsQueryParams{}
	items, _, _, err := lxService.Products(params)
	if err != nil {
		t.Errorf("lxService.Products error: %s", err.Error())
	} else {
		t.Log(jsonx.ToJson(items, "[]"))
	}
}

func TestService_Product(t *testing.T) {
	item, err := lxService.Product(0)
	if err != nil {
		t.Errorf("lxService.Product error: %s", err.Error())
	} else {
		t.Log(jsonx.ToJson(item, "[]"))
	}
}
