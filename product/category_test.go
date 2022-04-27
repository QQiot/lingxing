package product

import (
	"github.com/hiscaler/gox/jsonx"
	"testing"
)

func TestService_Categories(t *testing.T) {
	params := CategoriesQueryParams{}
	params.Limit = 1
	var categories []Category
	for {
		items, nextOffset, isLastPage, err := lxService.Categories(params)
		if err != nil {
			t.Errorf("lxService.Categories error: %s", err.Error())
		} else {
			categories = append(categories, items...)
		}
		if isLastPage {
			break
		}
		params.Offset = nextOffset
	}
	t.Log(jsonx.ToJson(categories, "[]"))
}
