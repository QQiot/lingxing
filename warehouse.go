package lingxing

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	jsoniter "github.com/json-iterator/go"
)

// 仓库
type warehouseService service

// Warehouse 仓库
type Warehouse struct {
	WID  int    `json:"wid"`  // 仓库 ID
	Name string `json:"name"` // 仓库名
	Type int    `json:"type"` // 仓库类型（1；本地、3：海外）
}

type WarehousesQueryParams struct {
	Paging
	Type    int `json:"type,omitempty"`     // 仓库类型（1；本地[默认]、3：海外）
	SubType int `json:"sub_type,omitempty"` // 海外仓子类型（1：无 API 海外仓、2：有 API 海外仓【此参数只有在type=3的时候有效】）
}

func (m WarehousesQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Type, validation.In(1, 3).Error("无效的仓库类型")),
		validation.Field(&m.SubType, validation.When(m.Type == 3, validation.In(1, 2).Error("无效的海外仓子类型"))),
	)
}

// All 查询本地仓库列表
// https://openapidoc.lingxing.com/#/docs/Warehouse/WarehouseLists
func (s warehouseService) All(params WarehousesQueryParams) (items []Warehouse, nextOffset int, isLastPage bool, err error) {
	params.SetPagingVars()
	if err = params.Validate(); err != nil {
		return
	}

	params.SetPagingVars()
	res := struct {
		NormalResponse
		Data []Warehouse `json:"data"`
	}{}
	resp, err := s.httpClient.R().
		SetBody(params).
		Post("/data/local_inventory/warehouse")
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
