package lingxing

import (
	"context"
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
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
	Paging
}

func (m CategoriesQueryParams) Validate() error {
	return nil
}

func (s productService) Categories(params CategoriesQueryParams) (items []Category, nextOffset int, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.SetPagingVars()
	res := struct {
		NormalResponse
		Data []Category `json:"data"`
	}{}
	resp, err := s.httpClient.R().
		SetBody(params).
		Post("/routing/data/local_inventory/category")
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

// 新增/更新分类
// https://openapidoc.lingxing.com/#/docs/Product/SetCategory

type UpsertCategory struct {
	ID        int    `json:"id"`         // 为零表示新增，不为零时表示编辑
	ParentCID int    `json:"parent_cid"` // 父级分类 ID
	Title     string `json:"title"`      // 分类名称
}

type UpsertCategoryRequest struct {
	Data []UpsertCategory `json:"data"`
}

func (m UpsertCategoryRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Data,
			validation.Required.Error("分类列表不能为空"),
			validation.Each(validation.WithContext(func(ctx context.Context, value interface{}) error {
				item, ok := value.(UpsertCategory)
				if !ok {
					return errors.New("无效的分类信息")
				}
				return validation.ValidateStruct(&item,
					validation.Field(&item.Title, validation.Required.Error("分类名称不能为空")),
				)
			})),
		),
	)
}

func (s productService) UpsertCategory(req UpsertCategoryRequest) (items []Category, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	res := struct {
		NormalResponse
		Data []UpsertCategory `json:"data"`
	}{}
	resp, err := s.httpClient.R().
		SetBody(req).
		Post("/routing/storage/category/set")
	if err != nil {
		return
	}

	if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
		items = make([]Category, len(res.Data))
		for i := range res.Data {
			items[i] = Category{
				CID:       res.Data[i].ID,
				ParentCID: res.Data[i].ParentCID,
				Title:     res.Data[i].Title,
			}
		}
	}

	return
}
