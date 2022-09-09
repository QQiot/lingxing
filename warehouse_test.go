package lingxing

import (
	"github.com/hiscaler/gox/jsonx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_warehouseService_All(t *testing.T) {
	tests := []struct {
		name     string
		params   WarehousesQueryParams
		hasError bool
	}{
		{"t0", WarehousesQueryParams{Type: 1234}, true},
		{"t1", WarehousesQueryParams{Type: 2, SubType: 1234}, true},
		{"t2", WarehousesQueryParams{}, false},
		{"t3", WarehousesQueryParams{Paging: Paging{Limit: 1}}, false},
		{"t4", WarehousesQueryParams{Paging: Paging{Limit: 1, Offset: 2}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			items, nextOffset, isLastPage, err := lingXingClient.Services.Warehouse.All(tt.params)
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

func Test_warehouseService_InboundOrders(t *testing.T) {
	tests := []struct {
		name     string
		params   InboundOrdersQueryParams
		hasError bool
	}{
		{"t0", InboundOrdersQueryParams{Type: 1234}, true},
		{"t1", InboundOrdersQueryParams{Type: 1, Status: 1}, true},
		{"t2", InboundOrdersQueryParams{}, true},
		{"t3", InboundOrdersQueryParams{
			Paging:          Paging{Limit: 1},
			SearchFieldTime: "create_time",
			OrderSn:         "123",
			WID:             "a",
			Type:            1,
			Status:          40,
			StartDate:       "2022-09-01",
			EndDate:         "2022-09-10",
		}, false},
		{"t4", InboundOrdersQueryParams{
			Paging:          Paging{Limit: 1, Offset: 2},
			SearchFieldTime: "create_time",
			OrderSn:         "123",
			Type:            1,
			WID:             "a",
			Status:          40,
			StartDate:       "2022-09-01",
			EndDate:         "2022-09-10",
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			items, nextOffset, isLastPage, err := lingXingClient.Services.Warehouse.InboundOrders(tt.params)
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
			} else {
				if !tt.hasError {
					t.Logf("error: %s", err.Error())
				}
				if n > 0 {
					assert.Equalf(t, 0, n, "All(%s) items count", params) // if error not equal nil, items will be an empty slice
				}
			}
		})
	}
}

func Test_warehouseService_OutboundOrders(t *testing.T) {
	tests := []struct {
		name     string
		params   OutboundOrdersQueryParams
		hasError bool
	}{
		{"t0", OutboundOrdersQueryParams{Type: 1234}, true},
		{"t1", OutboundOrdersQueryParams{Type: 1, Status: 1}, true},
		{"t2", OutboundOrdersQueryParams{}, true},
		{"t3", OutboundOrdersQueryParams{
			Paging:          Paging{Limit: 1},
			SearchFieldTime: "create_time",
			OrderSn:         "123",
			WID:             "a",
			Type:            11,
			Status:          40,
			StartDate:       "2022-09-01",
			EndDate:         "2022-09-10",
		}, false},
		{"t4", OutboundOrdersQueryParams{
			Paging:          Paging{Limit: 1, Offset: 2},
			SearchFieldTime: "create_time",
			OrderSn:         "123",
			Type:            11,
			WID:             "a",
			Status:          40,
			StartDate:       "2022-09-01",
			EndDate:         "2022-09-10",
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			items, nextOffset, isLastPage, err := lingXingClient.Services.Warehouse.OutboundOrders(tt.params)
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
			} else {
				if !tt.hasError {
					t.Logf("error: %s", err.Error())
				}
				if n > 0 {
					assert.Equalf(t, 0, n, "All(%s) items count", params) // if error not equal nil, items will be an empty slice
				}
			}
		})
	}
}
