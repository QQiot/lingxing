package basic

import (
	"encoding/json"
	"errors"
	"github.com/hiscaler/lingxing"
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
	SellerAccountId int    `json:"seller_account_id"` // SELLER_ID
	AccountName     string `json:"account_name"`
}

func (s service) Sellers() (items []Seller, err error) {
	res := struct {
		lingxing.NormalResponse
		Data []Seller `json:"data"`
	}{}
	resp, err := s.lingXing.Client.R().Get("/data/seller/lists")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
			if err = lingxing.ErrorWrap(res.Code, res.Message); err == nil {
				items = res.Data
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
