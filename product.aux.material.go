package lingxing

import jsoniter "github.com/json-iterator/go"

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
