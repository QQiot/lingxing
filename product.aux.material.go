package lingxing

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	jsoniter "github.com/json-iterator/go"
)

// 辅料
type productAuxMaterialService service

// 查询产品辅料列表
// https://openapidoc.lingxing.com/#/docs/Product/productAuxList

// ProductAuxMaterial 辅料
type ProductAuxMaterial struct {
	ID                    int                                 `json:"id"`                      // 辅料 ID
	SKU                   string                              `json:"sku"`                     // SKU
	ProductName           string                              `json:"product_name"`            // 品名
	CgPrice               float64                             `json:"cg_price"`                // 采购：采购价格（RMB）
	CgProductLength       float64                             `json:"cg_product_length"`       // 单品规格长(cm)
	CgProductWidth        float64                             `json:"cg_product_width"`        // 单品规格宽(cm)
	CgProductHeight       float64                             `json:"cg_product_height"`       // 单品规格高(cm)
	CgProductNetWeight    float64                             `json:"cg_product_net_weight"`   // 单品净重(g)
	PurchaseSupplierQuote []SupplierQuote                     `json:"purchase_supplier_quote"` // 供应商报价
	AuxRelationProduct    []ProductAuxMaterialRelationProduct `json:"aux_relation_product"`    // 关联产品
}

type ProductAuxMaterialRelationProduct struct {
	Pid         int    `json:"pid"`          // 产品id
	ProductName string `json:"product_name"` // 产品名称
	SKU         string `json:"sku"`          // SKU
	Quantity    int    `json:"quantity"`     // 关联辅料的数量
}

type ProductAuxMaterialsQueryParams struct {
	Paging
}

func (m ProductAuxMaterialsQueryParams) Validate() error {
	return nil
}

func (s productAuxMaterialService) All(params ProductAuxMaterialsQueryParams) (items []ProductAuxMaterial, nextOffset int, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.SetPagingVars()
	res := struct {
		NormalResponse
		Data []ProductAuxMaterial `json:"data"`
	}{}
	resp, err := s.httpClient.R().
		SetBody(params).
		Post("/routing/data/local_inventory/productAuxList")
	if err != nil {
		return
	}

	if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
		items = res.Data
		nextOffset = params.NextOffset
		isLastPage = len(items) < params.Limit
	}
	return
}

// 添加、编辑辅料
// https://openapidoc.lingxing.com/#/docs/Product/setAux

type UpsertProductAuxMaterialRequest struct {
	SKU                string                                  `json:"sku"`                             // SKU（添加时必填）
	ProductName        string                                  `json:"product_name"`                    // 品名（添加时必填）
	CgPrice            float64                                 `json:"cg_price,omitempty"`              // 采购：采购价格（RMB）
	CgProductLength    float64                                 `json:"cg_product_length,omitempty"`     // 单品规格长(cm)
	CgProductWidth     float64                                 `json:"cg_product_width,omitempty"`      // 单品规格宽(cm)
	CgProductHeight    float64                                 `json:"cg_product_height,omitempty"`     // 单品规格高(cm)
	CgProductNetWeight float64                                 `json:"cg_product_net_weight,omitempty"` // 单品净重(g)
	SupplierQuote      []UpsertProductAuxMaterialSupplierQuote `json:"supplier_quote,omitempty"`        // 采购商（2021.11.20号开始启用；不传该参数则清空产品供应商报价，否则覆盖
	Remark             string                                  `json:"remark"`                          // 辅料描述
}

// UpsertProductAuxMaterialSupplierQuote 采购商（2021.11.20号开始启用；不传该参数则清空产品供应商报价，否则覆盖
type UpsertProductAuxMaterialSupplierQuote struct {
	ERPSupplierId      string                                      `json:"erp_supplier_id"`      // 领星ERP供应商id
	SupplierId         string                                      `json:"supplier_id"`          // 客户系统供应商id，没有填这个值或者对应供应商不存在，则取erp_supplier_id
	SupplierProductURL []string                                    `json:"supplier_product_url"` // 采购链接，字符串数组，最多20个，没有则传空数组
	IsPrimary          int                                         `json:"is_primary"`           // 首选供应商(1:是；0：否)
	Quotes             []UpsertProductAuxMaterialSupplierQuoteItem `json:"quotes"`               // 报价信息
}

// UpsertProductAuxMaterialSupplierQuoteItem 报价信息
type UpsertProductAuxMaterialSupplierQuoteItem struct {
	Currency   string                                               `json:"currency"`    // 报价币种，目前只有CNY和USD
	IsTax      int                                                  `json:"is_tax"`      // 是否含税，0-否；1-是
	TaxRate    int                                                  `json:"tax_rate"`    // 税率，整数，0-99，为空则表示为0
	StepPrices []UpsertProductAuxMaterialSupplierQuoteItemStepPrice `json:"step_prices"` // 阶梯价信息
}

func (m UpsertProductAuxMaterialSupplierQuoteItem) Validte() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Currency,
			validation.Required.Error("报价币种不能为空"),
			validation.In("CNY", "USD").Error("无效的报价币种"),
		),
		validation.Field(&m.IsTax, validation.In(0, 1).Error("是否含税标识错误")),
		validation.Field(&m.TaxRate,
			validation.Min(0).Error("税率不能小于 {{.threshold}}"),
			validation.Max(99).Error("税率不能大于 {{.threshold}}"),
		),
	)
}

// UpsertProductAuxMaterialSupplierQuoteItemStepPrice  阶梯价信息
type UpsertProductAuxMaterialSupplierQuoteItemStepPrice struct {
	Moq          int     `json:"moq"`            // 最小起订量，最小值为1
	PriceWithTax float64 `json:"price_with_tax"` // 税单价，4位小数
}

func (m UpsertProductAuxMaterialRequest) Validate() error {
	return nil
}

func (s productAuxMaterialService) Upsert(req UpsertProductAuxMaterialRequest) (err error) {
	if err = req.Validate(); err != nil {
		return
	}

	_, err = s.httpClient.R().
		SetBody(req).
		Post("/routing/storage/product/setAux")
	return
}
