package lingxing

import (
	"errors"
	"github.com/hiscaler/gox/jsonx"
	"testing"
)

func TestProductService_Products(t *testing.T) {
	params := ProductsQueryParams{}
	params.Limit = 1
	var products []Product
	for {
		items, nextOffset, isLastPage, err := lingXingClient.Services.Product.Products(params)
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

func TestProductService_ProductNotFound(t *testing.T) {
	_, err := lingXingClient.Services.Product.Product(0)
	if !errors.Is(err, ErrNotFound) {
		t.Error("Services.Product.Products error is 'not found error'")
	}
}

func TestProductService_Product(t *testing.T) {
	item, err := lingXingClient.Services.Product.Product(63286)
	if err != nil {
		t.Errorf("Services.Product.Products error %s", err.Error())
	} else {
		t.Log(jsonx.ToPrettyJson(item))
	}
}
