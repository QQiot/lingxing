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
	Authorization   authorizationService
	BasicData       basicDataService
	CustomerService customerServiceService
	Product         productService
	Sale            saleService
	FBA             fbaService
}
