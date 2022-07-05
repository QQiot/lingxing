package lingxing

import (
	"errors"
	"github.com/hiscaler/gox/jsonx"
	"testing"
)

func TestProductServiceProduct_All(t *testing.T) {
	params := ProductsQueryParams{}
	params.Limit = 12
	var products []Product
	for {
		items, nextOffset, isLastPage, err := lingXingClient.Services.Product.All(params)
		if err != nil {
			t.Errorf("Services.Product.Products error: %s", err.Error())
		} else {
			products = append(products, items...)
		}
		if isLastPage || err != nil {
			break
		}
		params.Offset = nextOffset
	}
	t.Log(jsonx.ToPrettyJson(products))
}

func TestProductServiceProduct_OneNotFound(t *testing.T) {
	_, err := lingXingClient.Services.Product.One(0)
	if !errors.Is(err, ErrNotFound) {
		t.Error("Services.Product.Products error is 'not found error'")
	}
}

func TestProductServiceProduct_One(t *testing.T) {
	item, err := lingXingClient.Services.Product.One(63286)
	if err != nil {
		t.Errorf("Services.Product.Products error %s", err.Error())
	} else {
		t.Log(jsonx.ToPrettyJson(item))
	}
}
