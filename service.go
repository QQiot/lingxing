package lingxing

import (
	"github.com/go-resty/resty/v2"
	"github.com/hiscaler/lingxing/config"
	"log"
)

type service struct {
	config     *config.Config // Config
	logger     *log.Logger    // Logger
	httpClient *resty.Client  // HTTP client
}

// API Services
type services struct {
	Authorization   authorizationService   // 认证
	BasicData       basicDataService       // 基础数据
	CustomerService customerServiceService // 客服
	Product         productService         // 产品
	Sale            saleService            // 销售
	FBA             fbaService             // FBA
	Statistic       statisticService       // 统计
}
