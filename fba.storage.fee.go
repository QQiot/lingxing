package lingxing

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hiscaler/lingxing/constant"
	jsoniter "github.com/json-iterator/go"
)

// 仓储费
type fbaStorageFeeService service

type FBALongTermStorageFee struct {
	SID                                  int    `json:"sid"`                                     // 店铺ID
	SnapshotDate                         string `json:"snapshot_date"`                           // 时间
	SKU                                  string `json:"sku"`                                     // SKU
	FnSKU                                string `json:"fnsku"`                                   // FNSKU
	Asin                                 string `json:"asin"`                                    // ASIN
	ProductName                          string `json:"product_name"`                            // 标题
	Condition                            string `json:"condition"`                               // 状况
	QtyCharged12monthsLongTermStorageFee string `json:"qty_charged_12_mo_long_term_storage_fee"` // 12个月以上收费商品量
	PerUnitVolume                        string `json:"per_unit_volume"`                         // 单个商品体积
	Currency                             string `json:"currency"`                                // 币种
	TwelveMonthsLongTermsStorageFee      string `json:"12_mo_long_terms_storage_fee"`            // 12个月以上收费
	QtyCharged6MonthsLongTermStorageFee  string `json:"qty_charged_6_mo_long_term_storage_fee"`  // 6-12个月收费商品量
	SixMonthsLongTermsStorageFee         string `json:"6_mo_long_terms_storage_fee"`             // 6-12个月收费
	VolumeUnit                           string `json:"volume_unit"`                             // 体积单位
	Country                              string `json:"country"`                                 // 国家
	IsSmallAndLight                      string `json:"is_small_and_light"`
	EnrolledInSmallAndLight              string `json:"enrolled_in_small_and_light"`
}

type FBALongTermStorageFeesQueryParams struct {
	Paging
	SID       int    `json:"sid"`        // 店铺 ID
	StartDate string `json:"start_date"` // 收费日期左闭区间
	EndDate   string `json:"end_date"`   // 收费日期右开区间
}

func (m FBALongTermStorageFeesQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.SID, validation.Required.Error("店铺 ID 不能为空")),
		validation.Field(&m.StartDate,
			validation.Required.Error("收费开始日期不能为空"),
			validation.Date(constant.DateFormat).Error("收费开始日期格式有误，正确的格式为：2006-01-02"),
		),
		validation.Field(&m.EndDate,
			validation.Required.Error("收费结束日期不能为空"),
			validation.Date(constant.DateFormat).Error("收费结束日期格式有误，正确的格式为：2006-01-02"),
		),
	)
}

// LongTerm 查询 FBA 长期仓储费
func (s fbaStorageFeeService) LongTerm(params FBALongTermStorageFeesQueryParams) (items []FBALongTermStorageFee, nextOffset int, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.SetPagingVars()
	res := struct {
		NormalResponse
		Data []FBALongTermStorageFee `json:"data"`
	}{}
	resp, err := s.httpClient.R().
		SetBody(params).
		Post("/data/fba_report/storageFeeLongTerm")
	if err != nil {
		return
	}

	if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
		items = res.Data
		nextOffset = params.NextOffset
		isLastPage = len(items) < params.Limit
	}
	return
}

// 月仓储费

type FBAMonthStorageFee struct {
	SID                           int     `json:"sid"`                              // 店铺ID
	Asin                          string  `json:"asin"`                             // ASIN
	FnSKU                         string  `json:"fnsku"`                            // FNSKU
	ProductName                   string  `json:"product_name"`                     // 标题
	FulfillmentCenter             string  `json:"fulfillment_center"`               // 仓库编号
	CountryCode                   string  `json:"country_code"`                     // 国家代码
	LongestSide                   float64 `json:"longest_side"`                     // 长边
	MedianSide                    float64 `json:"median_side"`                      // 中间边
	ShortestSide                  float64 `json:"shortest_side"`                    // 短边
	MeasurementUnits              string  `json:"measurement_units"`                // 长中短边单位
	Weight                        float64 `json:"weight"`                           // 重量
	WeightUnits                   string  `json:"weight_units"`                     // 重量单位
	ItemVolume                    float64 `json:"item_volume"`                      // 体积
	VolumeUnits                   string  `json:"volume_units"`                     // 体积单位
	ProductSizeTier               string  `json:"product_size_tier"`                // 产品标准
	AverageQuantityOnHand         float64 `json:"average_quantity_on_hand"`         // 库存量
	AverageQuantityPendingRemoval float64 `json:"average_quantity_pending_removal"` // 待移除量
	EstimatedTotalItemVolume      float64 `json:"estimated_total_item_volume"`      // 总体积
	MonthOfCharge                 string  `json:"month_of_charge"`                  // 收费月份
	StorageRate                   float64 `json:"storage_rate"`                     // 收费标准
	Currency                      string  `json:"currency"`                         // 币种
	EstimatedMonthlyStorageFee    float64 `json:"estimated_monthly_storage_fee"`    // 预估仓储费
}

type FBAMonthStorageFeesQueryParams struct {
	Paging
	SID   int    `json:"sid"`   // 店铺 ID
	Month string `json:"month"` // 收费月份
}

func (m FBAMonthStorageFeesQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.SID, validation.Required.Error("店铺 ID 不能为空")),
		validation.Field(&m.Month,
			validation.Required.Error("收费月份不能为空"),
			validation.Date("2006-01").Error("收费月份格式有误，正确的格式为：2006-01"),
		),
	)
}

// Month FBA 月仓储费
func (s fbaStorageFeeService) Month(params FBAMonthStorageFeesQueryParams) (items []FBAMonthStorageFee, nextOffset int, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.SetPagingVars()
	res := struct {
		NormalResponse
		Data []FBAMonthStorageFee `json:"data"`
	}{}
	resp, err := s.httpClient.R().
		SetBody(params).
		Post("/data/fba_report/storageFeeMonth")
	if err != nil {
		return
	}

	if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
		items = res.Data
		nextOffset = params.NextOffset
		isLastPage = len(items) < params.Limit
	}
	return
}
