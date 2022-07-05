package lingxing

import (
	"fmt"
	"github.com/hiscaler/gox/jsonx"
	"testing"
)

func TestProductServiceCategory_All(t *testing.T) {
	params := CategoriesQueryParams{}
	params.Limit = 10
	var categories []Category
	for {
		items, nextOffset, isLastPage, err := lingXingClient.Services.Product.Category.All(params)
		if err != nil {
			t.Errorf("Services.Product.Category.All() error: %s", err.Error())
		} else {
			categories = append(categories, items...)
		}
		if isLastPage || err != nil {
			break
		}
		params.Offset = nextOffset
	}
	t.Log(jsonx.ToPrettyJson(categories))
}

func TestProductServiceCategory_Upsert(t *testing.T) {
	// Add
	req := UpsertCategoryRequest{
		Data: []UpsertCategory{
			{ID: 0, ParentCID: 0, Title: "PC"},
		},
	}
	brands, err := lingXingClient.Services.Product.Category.Upsert(req)
	if err != nil {
		t.Errorf("Services.Product.Category.Upsert() create error: %s", err.Error())
	} else {
		t.Log(jsonx.ToPrettyJson(brands))
		req.Data = make([]UpsertCategory, len(brands))
		for i, brand := range brands {
			req.Data[i].ID = brand.CID
			req.Data[i].Title = fmt.Sprintf("%s-%d", brand.Title, i)
			if _, err := lingXingClient.Services.Product.Category.Upsert(req); err != nil {
				t.Errorf("Services.Product.Category.Upsert() error: %s", err.Error())
			}
		}
	}
}
