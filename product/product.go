package product

import (
	"errors"
	"github.com/hiscaler/lingxing"
	jsoniter "github.com/json-iterator/go"
)

// 产品列表
// https://openapidoc.lingxing.com/#/docs/Product/ProductLists

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
	CgTransportCosts string          `json:"cg_transport_costs"` // 采购：运输成本
	CgPrice          string          `json:"cg_price"`           // 采购：采购价格（RMB）
	Status           int             `json:"status"`             // 状态编码
	StatusText       string          `json:"status_text"`        // 状态文本
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

// 产品详情
// https://openapidoc.lingxing.com/#/docs/Product/ProductDetails

type ProductPicture struct {
	PicURL    string `json:"pic_url"`    // 图片链接
	IsPrimary int    `json:"is_primary"` // 是否产品主图 0-否 1-是
}

// ProductLogistic 物流关联
type ProductLogistic struct {
	USCgTransportCosts string `json:"US_cg_transport_costs"` // 美国默认头程成本(含税)
	USCurrency         string `json:"US_currency"`           // 美国官方汇率code
	USBgImportHsCode   string `json:"US_bg_import_hs_code"`  // 报关：美国HS Code（进口国）
	USBgTaxRate        string `json:"US_bg_tax_rate"`        // 报关：美国税率
}

type ProductDetail struct {
	ID                       int               `json:"id"`                         // 产品id
	ProductName              string            `json:"product_name"`               // 产品名称
	SKU                      string            `json:"sku"`                        // SKU
	PicUrl                   string            `json:"pic_url"`                    // 上传的图片地址
	PictureList              []ProductPicture  `json:"picture_list"`               // 产品图片数组
	Model                    string            `json:"model"`                      // 产品型号
	Unit                     string            `json:"unit"`                       // 商品单位：套、个、台
	Status                   int               `json:"status"`                     // 状态：1-停售 2-在售 3-开发中 4-清仓
	CID                      int               `json:"cid"`                        // 分类ID
	BID                      int               `json:"bid"`                        // 品牌ID
	ProductDeveloper         string            `json:"product_developer"`          // 开发者
	ProductDeveloperUid      int               `json:"product_developer_uid"`      // 开发人
	Description              string            `json:"description"`                // 商品描述
	IsCombo                  int               `json:"is_combo"`                   // 1=组合商品；0=非组合商品
	Currency                 string            `json:"currency"`                   // 中国官方汇率code
	CgOptUsername            string            `json:"cg_opt_username"`            // 采购：采购员
	CgDelivery               int               `json:"cg_delivery"`                // 采购：交期
	CgPrice                  string            `json:"cg_price"`                   // 采购：采购价格（RMB）
	CgProductMaterial        string            `json:"cg_product_material"`        // 采购：材质
	CgProductLength          string            `json:"cg_product_length"`          // 采购：产品规格（CM）
	CgProductWidth           string            `json:"cg_product_width"`           // 采购：产品规格（CM）
	CgProductHeight          string            `json:"cg_product_height"`          // 采购：产品规格（CM）
	CgPackageLength          string            `json:"cg_package_length"`          // 采购：包装规格（CM）
	CgPackageWidth           string            `json:"cg_package_width"`           // 采购：包装规格（CM）
	CgPackageHeight          string            `json:"cg_package_height"`          // 采购：包装规格（CM）
	CgBoxLength              string            `json:"cg_box_length"`              // 采购：外箱规格（CM）
	CgBoxWidth               string            `json:"cg_box_width"`               // 采购：外箱规格（CM）
	CgBoxHeight              string            `json:"cg_box_height"`              // 采购：外箱规格（CM）
	CgProductNetWeight       string            `json:"cg_product_net_weight"`      // 采购：产品净重（G）
	CgProductGrossWeight     string            `json:"cg_product_gross_weight"`    // 采购：产品毛重（G）
	CgBoxWeight              string            `json:"cg_box_weight"`              // 采购：外箱实重（KG）
	CgBoxPcs                 int               `json:"cg_box_pcs"`                 // 采购：单箱数量（包装数量）
	BgCustomsExportName      string            `json:"bg_customs_export_name"`     // 报关：申报品名（中文）【中文报关名】
	BgCustomsImportName      string            `json:"bg_customs_import_name"`     // 报关：申报品名（英文）【英文报关名】
	BgCustomsImportPrice     string            `json:"bg_customs_import_price"`    // 报关：申报金额（进口国）【申报单价】
	BgExportHsCode           string            `json:"bg_export_hs_code"`          // 报关：HS Code（出口国）【中国HS Code】
	BgImportHsCode           string            `json:"bg_import_hs_code"`          // 报关：HS Code（进口国）【美国HS Code】
	BgTaxRate                string            `json:"bg_tax_rate"`                // 报关：税率【美国税率】
	BrandName                string            `json:"brand_name"`                 // 品牌名称
	CategoryName             string            `json:"category_name"`              // 分类名称
	SupplierQuote            []SupplierQuote   `json:"supplier_quote"`             // 供应商报价数据
	ComboProductList         []string          `json:"combo_product_list"`         // 组合商品列表
	ProductLogisticsRelation []ProductLogistic `json:"product_logistics_relation"` // 物流关联
}

func (s service) Product(id int) (item ProductDetail, err error) {
	res := struct {
		lingxing.NormalResponse
		Data ProductDetail `json:"data"`
	}{}
	resp, err := s.lingXing.Client.R().
		SetBody(map[string]int{"id": id}).
		Post("/routing/data/local_inventory/productInfo")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
			if err = lingxing.ErrorWrap(res.Code, res.Message); err == nil {
				item = res.Data
				if item.ID == 0 {
					err = lingxing.ErrNotFound
				}
			}
		}
	} else {
		if e := jsoniter.Unmarshal(resp.Body(), &res); e == nil {
			err = lingxing.ErrorWrap(res.Code, res.Message)
		} else {
			err = errors.New(resp.Status())
		}
	}
	return
}
