package lingxing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPurchaseService_Plans(t *testing.T) {
	params := PurchasePlansQueryParams{
		StartDate:            "2022-09-01",
		EndDate:              "2022-09-01",
		IsRelatedProcessPlan: false,
	}
	params.Limit = 2
	_, _, _, err := lingXingClient.Services.Purchase.Plans(params)
	assert.Equal(t, nil, err, "error")
}
