package product

import (
	"errors"
	"github.com/hiscaler/lingxing"
	jsoniter "github.com/json-iterator/go"
)

// SupplierQuote 供应商报价发
type SupplierQuote struct {
	PsqId              int                 `json:"psq_id"`
	ProductId          int                 `json:"product_id"`           // 产品ID
	SupplierId         int                 `json:"supplier_id"`          // 供应商ID
	IsPrimary          int                 `json:"is_primary"`           // 是否为首选供应商 0否 1是
	SupplierProductURL string              `json:"supplier_product_url"` // 采购链接
	SupplierName       string              `json:"supplier_name"`        // 供应商名称
	Quotes             []SupplierQuoteItem `json:"quotes"`               //	报价数据
}

// SupplierQuoteItem 供应商报价项
type SupplierQuoteItem struct {
	Currency     string                       `json:"currency"`      // 报价币种
	CurrencyIcon string                       `json:"currency_icon"` // 报价币种符号
	IsTax        int                          `json:"is_tax"`        // 是否含税 0否 1是
	TaxRate      int                          `json:"tax_rate"`      // 税率（百分比）0-99整数
	StepPrices   []SupplierQuoteItemStepPrice `json:"step_prices"`   // 报价梯度
}

type SupplierQuoteItemStepPrice struct {
	Moq          int     `json:"moq"`            // 最小采购量
	Price        float64 `json:"price"`          // 不含税单价
	PriceWithTax float64 `json:"price_with_tax"` // 含税单价
}

type Product struct {
	ID               int             `json:"id"`                 // 商品 ID
	CID              int             `json:"cid"`                // 类别ID
	CategoryName     string          `json:"category_name"`      // 类别ID
	BID              int             `json:"bid"`                // 品牌ID
	BrandName        string          `json:"brand_name"`         // 品牌
	SKU              string          `json:"sku"`                // SKU
	ProductName      string          `json:"product_name"`       // 品名
	PicURL           string          `json:"pic_url"`            // 图片链接
	CgDelivery       int             `json:"cg_delivery"`        // 采购：交期
	CgTransportCosts float64         `json:"cg_transport_costs"` // 采购：运输成本
	CgPrice          float64         `json:"cg_price"`           // 采购：采购价格（RMB）
	Status           int             `json:"status"`             // 状态编码
	StatusText       int             `json:"status_text"`        // 状态文本
	IsCombo          int             `json:"is_combo"`           // 是否为组合商品，0 = 否，1 = 是
	CreateTime       int             `json:"create_time"`        // 创建时间
	ProductDeveloper string          `json:"product_developer"`  // 开发人员
	CgOptUsername    string          `json:"cg_opt_username"`    // 采购：采购员
	SupplierQuote    []SupplierQuote `json:"supplier_quote"`     // 供应商报价
}

type ProductsQueryParams struct {
	lingxing.Paging
}

func (m ProductsQueryParams) Validate() error {
	return nil
}

func (s service) Products(params ProductsQueryParams) (items []Product, nextOffset int, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.SetPagingVars(params.Offset, params.Limit, s.lingXing.DefaultQueryParams.MaxLimit)
	res := struct {
		lingxing.NormalResponse
		Data []Product `json:"data"`
	}{}
	resp, err := s.lingXing.Client.R().
		SetBody(params).
		Post("/routing/data/local_inventory/productList")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
			if err = lingxing.ErrorWrap(res.Code, res.Message); err == nil {
				items = res.Data
				nextOffset = params.NextOffset
				isLastPage = res.Total <= params.Offset
			}
		}
	} else {
		if e := jsoniter.Unmarshal(resp.Body(), &res); e == nil {
			err = lingxing.ErrorWrap(res.Code, res.Message)
		} else {
			err = errors.New(resp.Status())
		}
	}
	if err != nil {
		return
	}

	return
}
