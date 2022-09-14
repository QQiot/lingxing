package lingxing

import (
	"context"
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	jsoniter "github.com/json-iterator/go"
)

// 本地产品品牌
// https://openapidoc.lingxing.com/#/docs/Product/Brand

type productBrandService service

type Brand struct {
	BID   int    `json:"bid"`   // 品牌 ID
	Title string `json:"title"` // 品牌名称
}

type BrandsQueryParams struct {
	Paging
}

func (m BrandsQueryParams) Validate() error {
	return nil
}

func (s productBrandService) All(params BrandsQueryParams) (items []Brand, nextOffset int, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.SetPagingVars()
	res := struct {
		NormalResponse
		Data []Brand `json:"data"`
	}{}
	resp, err := s.httpClient.R().
		SetBody(params).
		Post("/data/local_inventory/brand")
	if err != nil {
		return
	}

	if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
		items = res.Data
		nextOffset = params.nextOffset
		isLastPage = len(items) < params.Limit
	}
	return
}

// 新增/更新品牌
// https://openapidoc.lingxing.com/#/docs/Product/SetBrand

type UpsertBrand struct {
	ID    int    `json:"id"`    // 为零表示新增，不为零时表示编辑
	Title string `json:"title"` // 品牌名称
}

type UpsertBrandRequest struct {
	Data []UpsertBrand `json:"data"`
}

func (m UpsertBrandRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Data,
			validation.Required.Error("品牌列表不能为空"),
			validation.Each(validation.WithContext(func(ctx context.Context, value interface{}) error {
				item, ok := value.(UpsertBrand)
				if !ok {
					return errors.New("无效的品牌信息")
				}
				return validation.ValidateStruct(&item,
					validation.Field(&item.Title, validation.Required.Error("品牌名称不能为空")),
				)
			})),
		),
	)
}

func (s productBrandService) Upsert(req UpsertBrandRequest) (items []Brand, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	res := struct {
		NormalResponse
		Data []UpsertBrand `json:"data"`
	}{}
	resp, err := s.httpClient.R().
		SetBody(req).
		Post("/storage/brand/set")
	if err != nil {
		return
	}

	if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
		items = make([]Brand, len(res.Data))
		for i := range res.Data {
			items[i] = Brand{
				BID:   res.Data[i].ID,
				Title: res.Data[i].Title,
			}
		}
	}
	return
}
