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

// 亚马逊订单详情
// https://openapidoc.lingxing.com/#/docs/Sale/OrderDetail

type AmazonOrderDetailItem struct {
	SID                        string  `json:"sid"`                           // 店铺 ID
	Title                      string  `json:"title"`                         // 商品标题
	SellerSKU                  string  `json:"seller_sku"`                    // MSKU
	Asin                       string  `json:"asin"`                          // ASIN
	AsinURL                    string  `json:"asin_url"`                      // ASIN URL
	ProductId                  int     `json:"product_id"`                    // 本地商品ID
	SKU                        int     `json:"sku"`                           // 本地SKU
	ProductName                string  `json:"product_name"`                  // 品名
	PicURL                     string  `json:"pic_url"`                       // 图片链接
	OrderItemId                string  `json:"order_item_id"`                 // 图片链接
	UnitPriceAmount            float64 `json:"unit_price_amount"`             // 单价
	QuantityOrdered            int     `json:"quantity_ordered"`              // 下单量
	QuantityShipped            string  `json:"quantity_shipped"`              // 已配送
	SalesPriceAmount           float64 `json:"sales_price_amount"`            // 销售收益
	TaxAmount                  float64 `json:"tax_amount"`                    // 税费
	CgPrice                    float64 `json:"cg_price"`                      // 采购成本
	PromotionAmount            float64 `json:"promotion_amount"`              // 促消费
	CommissionAmount           float64 `json:"commission_amount"`             // 平台费
	OtherAmount                float64 `json:"other_amount"`                  // 亚马逊收取的其他费用，比如参与“Amazon Exlusives Program”产生的费用
	CgTransportCosts           float64 `json:"cg_transport_costs"`            // 头程费用
	FBAShipmentAmount          float64 `json:"fba_shipment_amount"`           // FBA发货费
	FeeName                    string  `json:"fee_name"`                      // 其他费名称，比如测评费
	FeeCost                    float64 `json:"fee_cost"`                      // 其他费金额，比如测评费
	FeeCurrency                string  `json:"fee_currency"`                  // 其他费币种，比如测评费
	FeeIcon                    string  `json:"fee_icon"`                      // 其他费币种符号，比如测评费
	Profit                     float64 `json:"profit"`                        //	毛利润
	ItemPriceAmount            float64 `json:"item_price_amount"`             //	商品支付金额
	ItemTaxAmount              float64 `json:"item_tax_amount"`               //	商品税
	ShippingPriceAmount        float64 `json:"shipping_price_amount"`         //	商品运费配送费
	ShippingTaxAmount          float64 `json:"shipping_tax_amount"`           //	商品运费税
	GiftWrapPriceAmount        float64 `json:"gift_wrap_price_amount"`        //	礼品包装费
	GiftWrapTaxAmount          float64 `json:"gift_wrap_tax_amount"`          // 礼品包装税
	ShippingDiscountAmount     float64 `json:"shipping_discount_amount"`      // 配送折扣
	ShippingDiscountTaxAmount  float64 `json:"shipping_discount_tax_amount"`  //	配送折扣税
	PromotionDiscountAmount    float64 `json:"promotion_discount_amount"`     //	商品促销折扣
	PromotionDiscountTaxAmount float64 `json:"promotion_discount_tax_amount"` //	商品促销折扣税
	CodFeeAmount               float64 `json:"cod_fee_amount"`                //	COD服务费用（货到付款服务费）
	CodFeeDiscountAmount       float64 `json:"cod_fee_discount_amount"`       //	COD服务费用折扣
}

type AmazonOrderDetail struct {
	SID               string              `json:"sid"`                 // 店铺 ID
	Country           string              `json:"country"`             // 国家（应平台要求，不再返回数据）
	City              string              `json:"city"`                // 城市（应平台要求，不再返回数据）
	District          string              `json:"district"`            // 地区（应平台要求，不再返回数据）
	OrderStatus       string              `json:"order_status"`        // 订单状态
	IsAssessed        int                 `json:"is_assessed"`         // 是否评测订单（1：是、0：否）
	OrderTotalAmount  float64             `json:"order_total_amount"`  // 订单总金额
	Currency          string              `json:"currency"`            // 订单金额币种
	Icon              string              `json:"icon"`                // 订单金额币种符号
	Phone             string              `json:"phone"`               // 手机号（应平台要求，不再返回数据）
	PostalCode        string              `json:"postal_code"`         // 邮编（应平台要求，不再返回数据）
	IsMcfOrder        int                 `json:"is_mcf_order"`        // 0普通订单,1多渠道订单
	IsBusinessOrder   int                 `json:"is_business_order"`   // 是否为B2B订单，1是，0不是
	CountryCode       string              `json:"country_code"`        // 国家代码（应平台要求，不再返回数据）
	PurchaseDateLocal string              `json:"purchase_date_local"` // 订购时间（站点时间）
	LastUpdateDate    string              `json:"last_update_date"`    // 订单更新站点时间
	ItemList          []AmazonOrderDetail `json:"item_list"`           // 订单明细
}

type AmazonOrderQueryParams struct {
	OrderId string `json:"order_id"` // 订单号
}

func (m AmazonOrderQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.OrderId, validation.Required.Error("订单号不能为空")),
	)
}

func (s service) AmazonOrder(params AmazonOrderQueryParams) (detail AmazonOrderDetail, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	res := struct {
		lingxing.NormalResponse
		Data AmazonOrderDetail `json:"data"`
	}{}
	resp, err := s.lingXing.Client.R().
		SetBody(params).
		Post("/data/mws/orderDetail")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
			detail = res.Data
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
