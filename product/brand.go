package product

import (
	"context"
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hiscaler/lingxing"
	jsoniter "github.com/json-iterator/go"
)

// 本地产品品牌
// https://openapidoc.lingxing.com/#/docs/Product/Brand

type Brand struct {
	BID   int    `json:"bid"`   // 品牌 ID
	Title string `json:"title"` // 品牌名称
}

type BrandsQueryParams struct {
	lingxing.Paging
}

func (m BrandsQueryParams) Validate() error {
	return nil
}

func (s service) Brands(params BrandsQueryParams) (items []Brand, nextOffset int, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.SetPagingVars(params.Offset, params.Limit, s.lingXing.DefaultQueryParams.MaxLimit)
	res := struct {
		lingxing.NormalResponse
		Data []Brand `json:"data"`
	}{}
	resp, err := s.lingXing.Client.R().
		SetBody(params).
		Post("/data/local_inventory/brand")
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

// 新增/更新品牌
// https://openapidoc.lingxing.com/#/docs/Product/SetBrand

type UpsertBrand struct {
	ID    int    `json:"id,omitempty"` // 为空时表新增，不为空时表编辑
	Title string `json:"title"`        // 品牌名称
}

type UpsertBrandRequest struct {
	Data []UpsertBrand `json:"data"`
}

func (m UpsertBrandRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Data,
			validation.Required.Error("品牌列表不能为空"),
			validation.Each(validation.WithContext(func(ctx context.Context, value interface{}) error {
				return validation.Validate(value, validation.Required.Error("品牌名称不能为空"))
			})),
		),
	)
}

func (s service) UpsertBrand(req UpsertBrandRequest) (items []Brand, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	res := struct {
		lingxing.NormalResponse
		Data []UpsertBrand `json:"data"`
	}{}
	resp, err := s.lingXing.Client.R().
		SetBody(req).
		Post("/storage/brand/set")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
			if err = lingxing.ErrorWrap(res.Code, res.Message); err == nil {
				for i := range res.Data {
					items = append(items, Brand{
						BID:   res.Data[i].ID,
						Title: res.Data[i].Title,
					})
				}
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
