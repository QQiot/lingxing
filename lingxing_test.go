package lingxing

import (
	"fmt"
	"github.com/hiscaler/lingxing/config"
	jsoniter "github.com/json-iterator/go"
	"os"
	"testing"
)

var lingXingClient *LingXing

func TestMain(m *testing.M) {
	b, err := os.ReadFile("./config/config_test.json")
	if err != nil {
		panic(fmt.Sprintf("Read config error: %s", err.Error()))
	}
	var c config.Config
	err = jsoniter.Unmarshal(b, &c)
	if err != nil {
		panic(fmt.Sprintf("Parse config file error: %s", err.Error()))
	}
	lingXingClient = NewLingXing(c)
	m.Run()
}
