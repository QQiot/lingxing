package product

import (
	"errors"
	"github.com/hiscaler/lingxing"
	jsoniter "github.com/json-iterator/go"
)

// 产品分类
// https://openapidoc.lingxing.com/#/docs/Product/Category

type Category struct {
	CID       int    `json:"cid"`        // 分类 ID
	ParentCID int    `json:"parent_cid"` // 父级分类 ID
	Title     string `json:"title"`      // 分类名称
}

type CategoriesQueryParams struct {
	lingxing.Paging
}

func (m CategoriesQueryParams) Validate() error {
	return nil
}

func (s service) Categories(params CategoriesQueryParams) (items []Category, nextOffset int, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.SetPagingVars(params.Offset, params.Limit, s.lingXing.DefaultQueryParams.MaxLimit)
	res := struct {
		lingxing.NormalResponse
		Data []Category `json:"data"`
	}{}
	resp, err := s.lingXing.Client.R().
		SetBody(params).
		Post("/routing/data/local_inventory/category")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
			if err = lingxing.ErrorWrap(res.Code, res.Message); err == nil {
				items = res.Data
				nextOffset = params.NextOffset
				isLastPage = len(items) < params.Limit
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
