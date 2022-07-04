package lingxing

import (
	"errors"
	"github.com/hiscaler/lingxing/entity"
	jsoniter "github.com/json-iterator/go"
)

// 查询亚马逊店铺信息
// https://openapidoc.lingxing.com/#/docs/BasicData/SellerLists

type basicDataService service

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

	if resp.IsSuccess() {
		if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
			if err = ErrorWrap(res.Code, res.Message); err == nil {
				items = res.Data
			}
		}
	} else {
		if e := jsoniter.Unmarshal(resp.Body(), &res); e == nil {
			err = ErrorWrap(res.Code, res.Message)
		} else {
			err = errors.New(resp.Status())
		}
	}
	return
}

// Accounts 查询ERP账号列表
func (s basicDataService) Accounts() (items []entity.Account, err error) {
	res := struct {
		NormalResponse
		Data []entity.Account `json:"data"`
	}{}
	resp, err := s.httpClient.R().Get("/data/account/lists")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
			if err = ErrorWrap(res.Code, res.Message); err == nil {
				items = res.Data
			}
		}
	} else {
		if e := jsoniter.Unmarshal(resp.Body(), &res); e == nil {
			err = ErrorWrap(res.Code, res.Message)
		} else {
			err = errors.New(resp.Status())
		}
	}
	return
}
