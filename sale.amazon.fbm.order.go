package lingxing

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hiscaler/lingxing/constant"
	jsoniter "github.com/json-iterator/go"
)

// 自发货订单
// https://openapidoc.lingxing.com/#/docs/Sale/FBMOrderList

type AmazonFBMOrder struct {
	OrderNumber           string   `json:"order_number"`            // 系统单号
	Status                string   `json:"status"`                  // 订单状态
	OrderFrom             string   `json:"order_from"`              // 订单类型
	CountryCode           string   `json:"country_code"`            // 目的国代码
	PurchaseTime          string   `json:"purchase_time"`           // 订购时间
	LogisticsTypeId       string   `json:"logistics_type_id"`       // 物流方式 ID
	LogisticsProviderId   string   `json:"logistics_provider_id"`   // 物流商 ID
	PlatformList          []string `json:"platform_list"`           // 平台订单号
	LogisticsTypeName     string   `json:"logistics_type_name"`     // 物流方式名称
	LogisticsProviderName string   `json:"logistics_provider_name"` // 物理商名称
	WID                   string   `json:"wid"`                     // 发货仓库 ID
	WarehouseName         string   `json:"warehouse_name"`          // 发货仓库名称
	CustomerComment       string   `json:"customer_comment"`        // 客服备注
}

type AmazonFBMOrdersQueryParams struct {
	Paging
	SID         string `json:"sid"`                    // 店铺 ID（多个使用逗号分隔开）
	OrderStatus string `json:"order_status,omitempty"` // 订单状态，用逗号分隔开（2:已发货、3:未付款、4:待审核、5:待发货、6:已取消）
	StartTime   string `json:"start_time,omitempty'"`  // 查询时间左闭区间，可精确到时分秒，格式：Y-m-d或Y-m-d H:i:s
	EndTime     string `json:"end_time,omitempty"`     // 查询时间右开区间，可精确到时分秒，格式：Y-m-d或Y-m-d H:i:s
}

func (m AmazonFBMOrdersQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.SID, validation.Required.Error("店铺 ID 不能为空")),
		validation.Field(&m.StartTime,
			validation.Required.Error("查询开始时间不能为空"),
			validation.Date(constant.DatetimeFormat).Error("查询开始时间格式有误"),
		),
		validation.Field(&m.EndTime,
			validation.Required.Error("查询结束时间不能为空"),
			validation.Date(constant.DatetimeFormat).Error("查询结束时间格式有误"),
		),
	)
}

// All 亚马逊自发货订单（FBM）列表
func (s fbmOrderService) All(params AmazonFBMOrdersQueryParams) (items []AmazonFBMOrder, nextOffset int, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.SetPagingVars()
	res := struct {
		NormalResponse
		Data []AmazonFBMOrder `json:"data"`
	}{}
	resp, err := s.httpClient.R().
		SetBody(params).
		Post("/routing/order/Order/getOrderList")
	if err != nil {
		return
	}

	if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
		items = res.Data
		nextOffset = params.NextOffset
		isLastPage = res.Total <= params.Offset
	}
	return
}
