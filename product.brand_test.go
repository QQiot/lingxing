package lingxing

import (
	"fmt"
	"github.com/hiscaler/gox/jsonx"
	"testing"
)

func TestProductService_Brands(t *testing.T) {
	params := BrandsQueryParams{}
	params.Limit = 10
	var brands []Brand
	for {
		items, nextOffset, isLastPage, err := lingXingClient.Services.Product.Brands(params)
		if err != nil {
			t.Errorf("Services.Product.Brands() error: %s", err.Error())
		} else {
			brands = append(brands, items...)
		}
		if isLastPage || err != nil {
			break
		}
		params.Offset = nextOffset
	}
	t.Log(jsonx.ToPrettyJson(brands))
}

func TestProductService_UpsertBrand(t *testing.T) {
	// Add
	req := UpsertBrandRequest{
		Data: []UpsertBrand{
			{ID: 0, Title: "HuaWei"},
		},
	}
	brands, err := lingXingClient.Services.Product.UpsertBrand(req)
	if err != nil {
		t.Errorf("Services.Product.UpsertBrand() create error: %s", err.Error())
	} else {
		t.Log(jsonx.ToPrettyJson(brands))
		req.Data = make([]UpsertBrand, len(brands))
		for i, brand := range brands {
			req.Data[i].ID = brand.BID
			req.Data[i].Title = fmt.Sprintf("%s-%d", brand.Title, i)
			if _, err := lingXingClient.Services.Product.UpsertBrand(req); err != nil {
				t.Errorf("Services.Product.UpsertBrand error: %s", err.Error())
			}
		}
	}
}
