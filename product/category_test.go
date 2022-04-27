package product

import (
	"fmt"
	"github.com/hiscaler/gox/jsonx"
	"testing"
)

func TestService_Categories(t *testing.T) {
	params := CategoriesQueryParams{}
	params.Limit = 10
	var categories []Category
	for {
		items, nextOffset, isLastPage, err := lxService.Categories(params)
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
	t.Log(jsonx.ToJson(categories, "[]"))
}

func TestService_UpsertCategory(t *testing.T) {
	// Add
	req := UpsertCategoryRequest{
		Data: []UpsertCategory{
			{ID: 0, ParentCID: 0, Title: "PC"},
		},
	}
	brands, err := lxService.UpsertCategory(req)
	if err != nil {
		t.Errorf("lxService.UpsertCategory create error: %s", err.Error())
	} else {
		t.Log(jsonx.ToJson(brands, "[]"))
		req.Data = make([]UpsertCategory, len(brands))
		for i, brand := range brands {
			req.Data[i].ID = brand.CID
			req.Data[i].Title = fmt.Sprintf("%s-%d", brand.Title, i)
			if _, err := lxService.UpsertCategory(req); err != nil {
				t.Errorf("lxService.UpsertCategory update error: %s", err.Error())
			}
		}
	}
}
