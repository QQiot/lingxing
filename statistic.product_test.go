package lingxing

import (
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
	_, _, _, err := lingXingClient.Services.Statistic.Products(params)
	assert.Equal(t, nil, err, "error")
}
