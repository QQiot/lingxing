package lingxing

import (
	"github.com/hiscaler/lingxing/entity"
	jsoniter "github.com/json-iterator/go"
)

// 查询亚马逊店铺信息
// https://openapidoc.lingxing.com/#/docs/BasicData/SellerLists

// Sellers 查询亚马逊店铺信息
func (s basicDataService) Sellers() (items []entity.Seller, err error) {
	res := struct {
		NormalResponse
		Data []entity.Seller `json:"data"`
	}{}
	resp, err := s.httpClient.R().Get("/data/seller/lists")
	if err != nil {
		return
	}

	if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
		items = res.Data
	}
	return
}
