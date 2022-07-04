package lingxing

import (
	"fmt"
	"github.com/hiscaler/gox/jsonx"
	"testing"
)

func TestProductService_Categories(t *testing.T) {
	params := CategoriesQueryParams{}
	params.Limit = 10
	var categories []Category
	for {
		items, nextOffset, isLastPage, err := lingXingClient.Services.Product.Categories(params)
		if err != nil {
			t.Errorf("lxService.Categories error: %s", err.Error())
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

func TestProductService_UpsertCategory(t *testing.T) {
	// Add
	req := UpsertCategoryRequest{
		Data: []UpsertCategory{
			{ID: 0, ParentCID: 0, Title: "PC"},
		},
	}
	brands, err := lingXingClient.Services.Product.UpsertCategory(req)
	if err != nil {
		t.Errorf("Services.Product.UpsertCategory() create error: %s", err.Error())
	} else {
		t.Log(jsonx.ToPrettyJson(brands))
		req.Data = make([]UpsertCategory, len(brands))
		for i, brand := range brands {
			req.Data[i].ID = brand.CID
			req.Data[i].Title = fmt.Sprintf("%s-%d", brand.Title, i)
			if _, err := lingXingClient.Services.Product.UpsertCategory(req); err != nil {
				t.Errorf("Services.Product.UpsertCategory() error: %s", err.Error())
			}
		}
	}
}
