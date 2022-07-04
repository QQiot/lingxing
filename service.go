package lingxing

import (
	"github.com/go-resty/resty/v2"
	"log"
)

type service struct {
	debug      bool          // Is debug mode
	logger     *log.Logger   // Logger
	httpClient *resty.Client // HTTP client
}

// API Services
type services struct {
	BasicData       basicDataService
	CustomerService customerServiceService
	Product         productService
	Sale            saleService
}
