package lingxing

import (
	jsoniter "github.com/json-iterator/go"
)

// 查询亚马逊店铺信息
// https://openapidoc.lingxing.com/#/docs/BasicData/SellerLists

type Seller struct {
	SID             int    `json:"sid"`               // 店铺 ID（领星ERP对企业已授权店铺的唯一标识）
	MID             int    `json:"mid"`               // 站点ID
	Name            string `json:"name"`              // 店铺名
	Country         string `json:"country"`           // 国家
	Region          string `json:"region"`            // 站点简称
	SellerId        string `json:"seller_id"`         // SELLER_ID
	SellerAccountId int    `json:"seller_account_id"` // 销售帐号 ID
	AccountName     string `json:"account_name"`      // 帐号名称
}

// Sellers 查询亚马逊店铺信息
func (s basicDataService) Sellers() (items []Seller, err error) {
	res := struct {
		NormalResponse
		Data []Seller `json:"data"`
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
