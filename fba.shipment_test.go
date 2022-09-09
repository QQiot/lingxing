package lingxing

import (
	"github.com/hiscaler/gox/jsonx"
	"github.com/stretchr/testify/assert"
	"testing"
)

// todo 内部错误，需联系领星
func Test_warehouseService_Plans(t *testing.T) {
	tests := []struct {
		name     string
		params   FBAShipmentPlansQueryParams
		hasError bool
	}{
		{"t0", FBAShipmentPlansQueryParams{}, true},
		{"t1", FBAShipmentPlansQueryParams{SearchFieldTime: "bad"}, true},
		{"t2", FBAShipmentPlansQueryParams{SearchFieldTime: "gmt_create"}, true},
		{"t3", FBAShipmentPlansQueryParams{
			Paging:          Paging{Limit: 1},
			SIDs:            "2345",
			WID:             "a",
			PackageType:     1,
			SearchFieldTime: "gmt_create",
			SearchField:     "order_sn",
			SearchValue:     "123",
			Status:          "1",
			MIDs:            "1",
			StartDate:       "2022-09-01",
			EndDate:         "2022-09-01",
		}, false},
		// {"t4", FBAShipmentPlansQueryParams{Paging: Paging{Limit: 1, Offset: 2}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			items, nextOffset, isLastPage, err := lingXingClient.Services.FBA.Shipment.Plans(tt.params)
			params := jsonx.ToJson(tt.params, "{}")
			assert.Equalf(t, tt.hasError, err != nil, "All(%s) error", params)
			n := len(items)
			if err == nil {
				if n > 0 {
					limit := tt.params.Limit
					if limit == 0 {
						limit = 1000 // Default size per page
					}
					assert.Equalf(t, true, n <= limit, "All(%s) items", params)                        // check return count is less or equal limit param value
					assert.Equalf(t, isLastPage, n < limit, "All(%s) isLastPage", params)              // check isLastPage value
					assert.Equalf(t, nextOffset, tt.params.Offset+limit, "All(%s) nextOffset", params) // check nextOffset value
				}
			} else if n > 0 {
				assert.Equalf(t, 0, n, "All(%s) items count", params) // if error not equal nil, items will be an empty slice
			}
		})
	}
}
