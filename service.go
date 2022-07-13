package lingxing

import (
	"github.com/go-resty/resty/v2"
	"log"
)

type cfg struct {
	debug     bool
	sandbox   bool
	appId     string
	appSecret string
}

type service struct {
	config     *cfg          // Config
	logger     *log.Logger   // Logger
	httpClient *resty.Client // HTTP client
}

// API Services
type services struct {
	Authorization   authorizationService
	BasicData       basicDataService
	CustomerService customerServiceService
	Product         productService
	Sale            saleService
	FBA             fbaService
}
