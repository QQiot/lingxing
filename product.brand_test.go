package lingxing

import (
	"fmt"
	"github.com/hiscaler/gox/jsonx"
	"testing"
)

func TestProductServiceBrand_All(t *testing.T) {
	params := BrandsQueryParams{}
	params.Limit = 10
	var brands []Brand
	for {
		items, nextOffset, isLastPage, err := lingXingClient.Services.Product.Brand.All(params)
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

func TestProductServiceBrand_Upsert(t *testing.T) {
	// Add
	req := UpsertBrandRequest{
		Data: []UpsertBrand{
			{ID: 0, Title: "HuaWei"},
		},
	}
	brands, err := lingXingClient.Services.Product.Brand.Upsert(req)
	if err != nil {
		t.Errorf("Services.Product.Brand.Upsert() create error: %s", err.Error())
	} else {
		t.Log(jsonx.ToPrettyJson(brands))
		req.Data = make([]UpsertBrand, len(brands))
		for i, brand := range brands {
			req.Data[i].ID = brand.BID
			req.Data[i].Title = fmt.Sprintf("%s-%d", brand.Title, i)
			if _, err := lingXingClient.Services.Product.Brand.Upsert(req); err != nil {
				t.Errorf("Services.Product.Brand.Upsert() error: %s", err.Error())
			}
		}
	}
}
