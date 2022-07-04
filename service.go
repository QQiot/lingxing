package lingxing

import (
	"github.com/go-resty/resty/v2"
	"log"
)

type service struct {
	debug              bool               // Is debug mode
	logger             *log.Logger        // Log
	httpClient         *resty.Client      // HTTP client
	defaultQueryParams defaultQueryParams // 查询默认值
}

// API Services
type services struct {
	BasicData       basicDataService
	CustomerService customerServiceService
	Product         productService
}
