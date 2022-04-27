package product

import (
	"github.com/hiscaler/gox/jsonx"
	"testing"
)

func TestService_Brands(t *testing.T) {
	params := BrandsQueryParams{}
	params.Limit = 1
	var brands []Brand
	for {
		items, nextOffset, isLastPage, err := lxService.Brands(params)
		if err != nil {
			t.Errorf("lxService.Products error: %s", err.Error())
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
