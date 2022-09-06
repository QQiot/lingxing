package lingxing

import (
	"fmt"
	"testing"
)

func TestStatisticService_Products(t *testing.T) {
	params := ProductStatisticQueryParams{
		StartDate: "2022-01-01",
		EndDate:   "2022-11-01",
	}
	items, _, _, err := lingXingClient.Services.Statistic.Products(params)
	_ = items
	fmt.Println(err)
}
