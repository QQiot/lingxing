package sale

import (
	"encoding/json"
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hiscaler/lingxing"
	"github.com/hiscaler/lingxing/constant"
	jsoniter "github.com/json-iterator/go"
	"time"
)

// 亚马逊订单
// https://openapidoc.lingxing.com/#/docs/Sale/Orderlists

const (
	DateTypeSaleTime        = 1 // 下单日期
	DateTypeOrderUpdateTime = 2 // 订单更新时间
)

// AmazonOrderItem 商品列表
type AmazonOrderItem struct {
	ASIN            string `json:"asin"`             // ASIN
	QuantityOrdered int    `json:"quantity_ordered"` // 数量
	SellerSKU       string `json:"seller_sku"`       // 卖家 SKU
}

type AmazonOrder struct {
	AmazonOrderId          string            `json:"amazon_order_id"`           // 订单号
	PurchaseDateLocal      time.Time         `json:"purchase_date_local"`       // 下单时间
	OrderStatus            string            `json:"order_status"`              // 订单状态
	OrderTotalCurrencyCode string            `json:"order_total_currency_code"` // 币种
	OrderTotalAmount       float64           `json:"order_total_amount"`        // 订单金额
	FulfillmentChannel     string            `json:"fulfillment_channel"`       // 发货渠道（亚马逊订单：AFN，自发货：MFN）
	BuyerEmail             string            `json:"buyer_email"`               // 买家邮件（应平台要求，不再返回数据）
	IsReturn               int               `json:"is_return"`                 // 是否退款（0：未退款、1：退款中、2：退款完成）
	IsMcfOrder             bool              `json:"is_mcf_order"`              // 是否多渠道订单（1：是、0：否）
	IsAssessed             bool              `json:"is_assessed"`               // 是否评测订单（1：是、0：否）
	EarliestShipDate       time.Time         `json:"earliest_ship_date"`        // 发货时限（2020-11-02T08:00:00Z）
	ShipmentDate           time.Time         `json:"shipment_date"`             // 发货日期
	LastUpdateDate         time.Time         `json:"last_update_date"`          // 订单更新站点时间
	SellerName             string            `json:"seller_name"`               // 店铺名称
	TrackingNumber         string            `json:"tracking_number"`           // 物流运单号
	PostalCode             string            `json:"postal_code"`               // 邮编（应平台要求，不再返回数据）
	Phone                  string            `json:"phone"`                     // 电话（应平台要求，不再返回数据）
	PostedDate             time.Time         `json:"posted_date"`               // 付款时间
	ItemList               []AmazonOrderItem `json:"item_list"`                 // 商品列表
}

type AmazonOrdersQueryParams struct {
	lingxing.Paging
	SID       int    `json:"sid"`                 // 店铺 ID
	StartDate string `json:"start_date"`          // 查询时间左闭区间，可精确到时分秒，格式：Y-m-d或Y-m-d H:i:s
	EndDate   string `json:"end_date"`            // 查询时间右开区间，可精确到时分秒，格式：Y-m-d或Y-m-d H:i:s
	DateType  int    `json:"date_type,omitempty"` // 日期类型，1：下单日期，2：订单更新时间，不填默认1
}

func (m AmazonOrdersQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.SID, validation.Required.Error("店铺 ID 不能为空")),
		validation.Field(&m.StartDate,
			validation.Required.Error("查询开始时间不能为空"),
			validation.Date(constant.DatetimeFormat).Error("查询开始时间格式有误"),
		),
		validation.Field(&m.EndDate,
			validation.Required.Error("查询结束时间不能为空"),
			validation.Date(constant.DatetimeFormat).Error("查询结束时间格式有误"),
		),
		validation.Field(&m.DateType, validation.In(DateTypeSaleTime, DateTypeOrderUpdateTime).Error("无效的日期类型")),
	)
}

func (s service) AmazonOrders(params AmazonOrdersQueryParams) (items []AmazonOrder, nextOffset int, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.SetPagingVars(params.Offset, params.Limit, s.lingXing.DefaultQueryParams.MaxLimit)
	res := struct {
		lingxing.NormalResponse
		Data []AmazonOrder `json:"data"`
	}{}
	resp, err := s.lingXing.Client.R().
		SetBody(params).
		Post("/data/mws/orders")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
			items = res.Data
			nextOffset = params.NextOffset
			isLastPage = res.Total <= params.Offset
		}
	} else {
		if e := json.Unmarshal(resp.Body(), &res); e == nil {
			err = lingxing.ErrorWrap(res.Code, res.Message)
		} else {
			err = errors.New(resp.Status())
		}
	}
	if err != nil {
		return
	}

	return
}
