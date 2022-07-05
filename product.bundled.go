package lingxing

import (
	jsoniter "github.com/json-iterator/go"
)

// 捆绑产品
type productBundledService service

// 查询捆绑产品关系列表
// https://openapidoc.lingxing.com/#/docs/Product/bundledProductList

// BundledProduct 捆绑产品
type BundledProduct struct {
	ID              int                  `json:"id"`               // 捆绑产品 ID
	SKU             string               `json:"sku"`              // 捆绑产品 SKU
	ProductName     string               `json:"product_name"`     // 捆绑产品名
	CgPrice         float64              `json:"cg_price"`         // 捆绑产品采购成本
	StatusText      string               `json:"status_text"`      // 产品状态
	BundledProducts []BundledProductItem `json:"bundled_products"` // 捆绑产品
}

// BundledProductItem 捆绑产品
type BundledProductItem struct {
	productId string `json:"productId"`  // 子产品ID
	SKU       string `json:"sku"`        // 子产品SKU
	Quantity  int    `json:"bundledQty"` // 捆绑数量
}

type BundledProductsQueryParams struct {
	Paging
}

func (m BundledProductsQueryParams) Validate() error {
	return nil
}

func (s productBundledService) All(params BundledProductsQueryParams) (items []BundledProduct, nextOffset int, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.SetPagingVars()
	res := struct {
		NormalResponse
		Data []BundledProduct `json:"data"`
	}{}
	resp, err := s.httpClient.R().
		SetBody(params).
		Post("/routing/data/local_inventory/bundledProductList")
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
