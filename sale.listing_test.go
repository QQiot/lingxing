package lingxing

import (
	"github.com/hiscaler/gox/jsonx"
	"testing"
)

func TestListingService_All(t *testing.T) {
	params := ListingsQueryParams{
		SID: 172,
	}
	items, _, _, err := lingXingClient.Services.Sale.Listing.All(params)
	if err != nil {
		t.Errorf("Services.Sale.Listing.All() error: %s", err.Error())
	} else {
		t.Log(jsonx.ToPrettyJson(items))
	}
}
