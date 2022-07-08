package lingxing

import (
	"github.com/hiscaler/gox/jsonx"
	"testing"
)

func TestCustomerServiceReviewService_All(t *testing.T) {
	params := CustomerServiceReviewsQueryParams{
		StartDate: "2022-01-01",
		EndDate:   "2022-11-01",
	}
	params.Limit = 200
	var reviews []CustomerServiceReview
	for {
		items, nextOffset, isLastPage, err := lingXingClient.Services.CustomerService.Review.All(params)
		if err != nil {
			t.Errorf("Services.CustomerService.Review.All() error: %s", err.Error())
		} else {
			reviews = append(reviews, items...)
		}
		if isLastPage || err != nil {
			break
		}
		params.Offset = nextOffset
	}
	t.Log(jsonx.ToPrettyJson(reviews))
}
