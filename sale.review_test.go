package lingxing

import (
	"github.com/hiscaler/gox/jsonx"
	"testing"
)

func TestReviewService_All(t *testing.T) {
	params := ReviewsQueryParams{
		SID:       172,
		StartDate: "2021-01-01",
		EndDate:   "2022-12-01",
	}
	items, _, _, err := lingXingClient.Services.Sale.Review.All(params)
	if err != nil {
		t.Errorf("Services.Sale.Review.All() error: %s", err.Error())
	} else {
		t.Log(jsonx.ToPrettyJson(items))
	}
}
