package lingxing

import (
	"fmt"
	"github.com/hiscaler/gox/jsonx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStatisticService_Products(t *testing.T) {
	params := ProductStatisticQueryParams{
		SID:       2345,
		StartDate: "2022-09-01",
		EndDate:   "2022-09-01",
	}
	params.Limit = 1
	items, _, _, err := lingXingClient.Services.Statistic.Products(params)
	assert.Equal(t, nil, err, "error")
	fmt.Println(jsonx.ToPrettyJson(items))
}
