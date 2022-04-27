package product

import (
	"fmt"
	"github.com/hiscaler/gox/jsonx"
	"testing"
)

func TestService_Brands(t *testing.T) {
	params := BrandsQueryParams{}
	params.Limit = 11
	var brands []Brand
	for {
		items, nextOffset, isLastPage, err := lxService.Brands(params)
		if err != nil {
			t.Errorf("lxService.Brands error: %s", err.Error())
		} else {
			brands = append(brands, items...)
		}
		if isLastPage {
			break
		}
		params.Offset = nextOffset
	}
	t.Log(jsonx.ToJson(brands, "[]"))
}

func TestService_UpsertBrand(t *testing.T) {
	// Add
	req := UpsertBrandRequest{
		Data: []UpsertBrand{
			{ID: 0, Title: "HuaWei"},
		},
	}
	brands, err := lxService.UpsertBrand(req)
	if err != nil {
		t.Errorf("lxService.UpsertBrand create error: %s", err.Error())
	} else {
		t.Log(jsonx.ToJson(brands, "[]"))
		req.Data = make([]UpsertBrand, len(brands))
		for i, brand := range brands {
			req.Data[i].ID = brand.BID
			req.Data[i].Title = fmt.Sprintf("%s-%d", brand.Title, i)
			if _, err := lxService.UpsertBrand(req); err != nil {
				t.Errorf("lxService.UpsertBrand update error: %s", err.Error())
			}
		}
	}
}
