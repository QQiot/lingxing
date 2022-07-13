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

	b, err := os.ReadFile("./config/config.json")
	if err != nil {
		panic(fmt.Sprintf("Read config error: %s", err.Error()))
	}
	c := struct {
		Debug       bool
		Sandbox     bool
		Environment struct {
			Dev struct {
				AppId     string
				AppSecret string
			}
			Prod struct {
				AppId     string
				AppSecret string
			}
		}
	}{}
	err = jsoniter.Unmarshal(b, &c)
	if err != nil {
		panic(fmt.Sprintf("Parse config file error: %s", err.Error()))
	}

	cfg := config.Config{
		Debug:   c.Debug,
		Sandbox: c.Sandbox,
	}
	if c.Sandbox {
		cfg.AppId = c.Environment.Dev.AppId
		cfg.AppSecret = c.Environment.Dev.AppSecret
	} else {
		cfg.AppId = c.Environment.Prod.AppId
		cfg.AppSecret = c.Environment.Prod.AppSecret
	}
	lingXingClient = NewLingXing(cfg)
	m.Run()
}
