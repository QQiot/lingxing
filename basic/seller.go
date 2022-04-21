package basic

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hiscaler/lingxing"
)

// 查询亚马逊店铺信息
// https://openapidoc.lingxing.com/#/docs/BasicData/SellerLists

type Seller struct {
	SID      string `json:"sid"`       // 店铺 ID（领星ERP对企业已授权店铺的唯一标识）
	MID      string `json:"mid"`       // 站点ID
	Name     string `json:"name"`      // 店铺名
	Country  string `json:"country"`   // 国家
	Region   string `json:"region"`    // 站点简称
	SellerId string `json:"seller_id"` // SELLER_ID
}

func (s service) Sellers() (items []Seller, err error) {
	res := struct {
		lingxing.NormalResponse
		Data []Seller `json:"data"`
	}{}
	resp, err := s.lingXing.Client.R().
		Post("/data/seller/lists")
	if err != nil {
		fmt.Println("ddd", err.Error())
		return
	}

	if resp.IsSuccess() {
		err = lingxing.ErrorWrap(res.Code, res.Message)
		if err == nil {
			json.Unmarshal(resp.Body(), &res)
			// res  = resp.RawResponse.
			// items = res.Data
		}
	} else {
		if e := json.Unmarshal(resp.Body(), &res); e == nil {
			err = lingxing.ErrorWrap(res.Code, res.Message)
		} else {
			err = errors.New(resp.Status())
		}
	}
	return
}
