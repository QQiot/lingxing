package lingxing

import (
	"context"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hiscaler/lingxing/constant"
	jsoniter "github.com/json-iterator/go"
	"time"
)

// https://openapidoc.lingxing.com/#/docs/MultiPlatform/MultiPlatOrder

type multiPlatformOrderService service

// 多平台订单

// MultiPlatformOrderAddress 收件信息
type MultiPlatformOrderAddress struct {
	AddressLine1        string `json:"address_line_1"`        // 详细地址 1
	AddressLine2        string `json:"address_line_2"`        // 详细地址 2
	AddressLine3        string `json:"address_line_3"`        // 详细地址 3
	City                string `json:"city"`                  // 城市
	District            string `json:"district"`              // 区县
	DoorplateNo         string `json:"doorplate_no"`          // 门牌号
	PostalCode          string `json:"postal_code"`           // 邮编
	ReceiverCountryCode string `json:"receiver_country_code"` // 目的国家二位简码
	ReceiverMobile      string `json:"receiver_mobile"`       // 收件人手机号
	ReceiverName        string `json:"receiver_name"`         // 收件人姓名
	ReceiverTel         string `json:"receiver_tel"`          // 收件人电话
	StateOrRegion       string `json:"state_or_region"`       // 收件人电话
}

// MultiPlatformOrderBuyer 买家信息
type MultiPlatformOrderBuyer struct {
	BuyerEmail string `json:"buyer_email"` // 买家邮箱
	BuyerName  string `json:"buyer_name"`  // 买家姓名
	BuyerNo    string `json:"buyer_no"`    // 平台买家 ID
	BuyerNote  string `json:"buyer_note"`  // 买家备注
}

// MultiPlatformOrderItem 商品信息
type MultiPlatformOrderItem struct {
	ItemFromName           string `json:"item_from_name"`           // 商品来源（1：线上、2：本地、3：本地(自定义平台导入带MSKU商品，可换货)）
	LocalProductName       string `json:"local_product_name"`       // 品名
	CustomerShippingAmount string `json:"customer_shipping_amount"` // 客付运费（币种取 amount_currency）
	CustomerTipAmount      string `json:"customer_tip_amount"`      // Shopify 小费（币种取 amount_currency）
	DiscountAmount         string `json:"discount_amount"`          // 折扣（币种取 amount_currency）
	ItemPriceAmount        string `json:"item_price_amount"`        // 商品金额（币种取 amount_currency）
	LocalSKU               string `json:"local_sku"`                // SKU
	MSKU                   string `json:"msku"`                     // MSKU
	OrderItemNo            string `json:"order_item_no"`            // 订单明细单号
	PlatformOrderNo        string `json:"platform_order_no"`        // 平台单号（一个系统单合并订单情况下会存在多个平台单号）
	PlatformStatus         string `json:"platform_status"`          // 平台订单商品状态
	Quantity               string `json:"quantity"`                 // 数量
	Remark                 string `json:"remark"`                   // 备注
	CustomizedUrl          string `json:"customized_url"`           // 亚马逊定制商品文件下载链接
	StockCost              string `json:"stockCost"`                // 商品出库成本
	TaxAmount              string `json:"tax_amount"`               // 商品税费（币种取 amount_currency）
	TransactionFeeAmount   string `json:"transaction_fee_amount"`   // 商品交易费（币种取 amount_currency）
	Type                   string `json:"type"`                     // 商品类型
	UnitPriceAmount        string `json:"unit_price_amount"`        // 单价（币种取 amount_currency）
	VariantAttr            string `json:"variant_attr"`             // 	变体属性
}

// MultiPlatformOrderLogistics 物流信息
type MultiPlatformOrderLogistics struct {
	ActualCarrier         string  `json:"actual_carrier"`          // 实际承运人
	CostAmount            int     `json:"cost_amount"`             // 物流实际运费
	CostCurrencyCode      string  `json:"cost_currency_code"`      // 物流实际运费币种 code
	LogisticsProviderId   int     `json:"logistics_provider_id"`   // 物流商 ID
	LogisticsProviderName string  `json:"logistics_provider_name"` // 物流商名称
	LogisticsTime         int     `json:"logistics_time"`          // 物流下单成功时间
	LogisticsTypeId       int     `json:"logistics_type_id"`       // 物流方式 ID
	LogisticsTypeName     int     `json:"logistics_type_name"`     // 物流方式名称
	PkgFeeWeight          float64 `json:"pkg_fee_weight"`          // 实际计费重量
	PkgFeeWeightUnit      string  `json:"pkg_fee_weight_unit"`     // 实际计费重单位
	PkgHeight             float64 `json:"pkg_height"`              // 实际包裹高
	PkgLength             float64 `json:"pkg_length"`              // 实际包裹长
	PkgSizeUnit           string  `json:"pkg_size_unit"`           // 包裹尺寸单位
	PkgWidth              float64 `json:"pkg_width"`               // 实际包裹宽
	PreCostAmount         int     `json:"pre_cost_amount"`         // 预估运费 负
	PreFeeWeight          float64 `json:"pre_fee_weight"`          // 预估计费重量
	PreFeeWeightUnit      string  `json:"pre_fee_weight_unit"`     // 预估计费重单位
	PrePkgHeight          float64 `json:"pre_pkg_height"`          // 预估包裹高（cm）
	PrePkgLength          float64 `json:"pre_pkg_length"`          // 预估包裹长（cm）
	PrePkgWidth           float64 `json:"pre_pkg_width"`           // 预估包裹宽（cm）
	PreWeight             float64 `json:"pre_weight"`              // 预估重量（g）
	Status                string  `json:"status"`                  // 状态（0：待物流下单、1：物流下单中、2：成功、3：失败、4：已取消）
}

// 	MultiPlatformOrderTag 标签+处理类型
type MultiPlatformOrderTag struct {
	TagName string `json:"tag_name"` // 标签名称
	TagNo   string `json:"tag_no"`   // 标签 NO
	TagType string `json:"tag_type"` // 标签类型（1：自定义处理类型、2：自定义订单标签、3：系统处理类型、4：系统订单标签）

}

// 	MultiPlatformOrderPlatformInfo 平台单信息
type MultiPlatformOrderPlatformInfo struct {
	CancelTime        int    `json:"cancel_time"`         // 取消单时间
	DeliveryTime      int    `json:"delivery_time"`       // 平台发货时间
	LatestShipTime    int    `json:"latest_ship_time"`    // 最后发货时间
	OrderFrom         string `json:"order_from"`          // 订单来源
	PaymentStatus     string `json:"payment_status"`      // 平台单支付状态
	PaymentTime       int    `json:"payment_time"`        // 支付时间
	PlatformCode      string `json:"platform_code"`       // 平台 CODE
	PlatformOrderName string `json:"platform_order_name"` // 平台订单别名 name
	PlatformOrderNo   string `json:"platform_order_no"`   // 平台订单号
	PurchaseTime      int    `json:"purchase_time"`       // 订购时间
	ShippingStatus    string `json:"shipping_status"`     // 平台单发货状态
	Status            string `json:"status"`              // 平台单状态
	StoreCountryCode  string `json:"store_Country_code"`  // 订单国家(二位iso_3166_1)
}

// 	MultiPlatformOrderTransactionInfo 交易信息
type MultiPlatformOrderTransactionInfo struct {
	CustomerShippingAmount float64 `json:"customer_shipping_amount"` // 客付运费（币种取 amount_currency）
	CustomerTaxAmountShow  float64 `json:"customer_tax_amount_show"` // 客付税费（币种取 amount_currency）
	CustomerTipAmount      float64 `json:"customer_tip_amount"`      // 小费（币种取 amount_currency）
	DiscountAmount         float64 `json:"discount_amount"`          // 折扣（币种取 amount_currency）
	OrderItemAmount        float64 `json:"order_item_amount"`        // 商品金额（币种取 amount_currency）
	OrderTotalAmount       float64 `json:"order_total_amount"`       // 订单总额（币种取 amount_currency）
	OutboundCostAmount     float64 `json:"outbound_cost_amount"`     // 预估出库成本（币种默认 CNY）
	PreCostAmount          float64 `json:"pre_cost_amount"`          // 预估运费（币种默认 CNY）
	ProfitAmount           float64 `json:"profit_amount"`            // 预估毛利润（币种默认 CNY）
	TransactionFeeAmount   float64 `json:"transaction_fee_amount"`   // 交易费（币种默认 CNY）
}

type MultiPlatformOrder struct {
	ID                     int                               `json:"id"`                       // ID
	AddressInfo            MultiPlatformOrderAddress         `json:"address_info"`             // 收件信息
	AmountCurrency         string                            `json:"amount_currency"`          // 币种
	BuyersInfo             MultiPlatformOrderBuyer           `json:"buyers_info"`              // 买家信息
	DeliveryType           string                            `json:"delivery_type"`            // 发货方式（自发货、平台发货【指由平台仓库自动完成履约的订单，如Walmart的WFS订单】）
	GlobalCancelTime       int                               `json:"global_cancel_time"`       // 取消时间
	GlobalDeliveryTime     int                               `json:"global_delivery_time"`     // 发货时间
	GlobalDistributionTime int                               `json:"global_distribution_time"` // 配货时间
	GlobalLatestShipTime   int                               `json:"global_latest_ship_time"`  // 发货时限
	GlobalOrderNo          string                            `json:"global_order_no"`          // 系统单号（系统本地唯一订单号）
	GlobalPaymentTime      int                               `json:"global_payment_time"`      // 付款时间
	GlobalPrintTime        int                               `json:"global_print_time"`        // 打单时间
	GlobalPurchaseTime     int                               `json:"global_purchase_time"`     // 订购时间
	GlobalReviewTime       int                               `json:"global_review_time"`       // 审核时间
	ItemInfo               []MultiPlatformOrderItem          `json:"item_info"`                // 商品信息
	LogisticsInfo          []MultiPlatformOrderLogistics     `json:"logistics_info"`           // 物流信息
	OrderFromName          string                            `json:"order_from_name"`          // 订单来源
	OrderStatus            string                            `json:"order_status"`             // 系统订单状态（1：同步中、2：已同步、3：未付款、4：待审核、5：待发货、6：已发货、7：已取消、8：不显示、9：平台发货）
	OrderTag               MultiPlatformOrderTag             `json:"order_tag"`                // 标签+处理类型
	OriginalGlobalOrderNo  string                            `json:"original_global_order_no"` // 补发订单原系统单号
	PlatformInfo           MultiPlatformOrderPlatformInfo    `json:"platform_info"`            // 平台单信息
	PlatformOrderName      string                            `json:"platform_order_name"`      // 平台名称
	PlatformOrderNo        string                            `json:"platform_order_no"`        // 平台 ID
	Remark                 string                            `json:"remark"`                   // 备注
	SplitType              string                            `json:"split_type"`               // 拆分单类型（1：原始单、2：合并单、3：拆分单）
	StoreId                string                            `json:"store_id"`                 // 店铺 ID
	TransactionInfo        MultiPlatformOrderTransactionInfo `json:"transaction_info"`         // 交易信息
	UpdateTime             int                               `json:"update_time"`              // 订单更新时间
	Wid                    string                            `json:"wid"`                      // 仓库ID
}

type MultiPlatformOrdersQueryParams struct {
	Paging
	DateType     string   `json:"date_type,omitempty"`     // 时间类型（更新时间：update_time、订购时间：global_purchase_time、发货时间：global_delivery_time）
	PlatformCode []string `json:"platform_code,omitempty"` // 平台 Code
	StartTime    string   `json:"start_time"`              // 开始时间（Y-m-d H:i:s 格式）
	EndTime      string   `json:"end_time"`                // 结束时间（Y-m-d H:i:s 格式）
	StoreId      []string `json:"store_id,omitempty"`      // 店铺 ID
}

func (m MultiPlatformOrdersQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.DateType, validation.When(m.DateType != "", validation.In("update_time", "global_purchase_time", "global_delivery_time").Error("无效的时间类型"))),
		validation.Field(&m.PlatformCode, validation.When(len(m.PlatformCode) > 0, validation.Each(validation.WithContext(func(ctx context.Context, value interface{}) error {
			code, ok := value.(string)
			if !ok {
				return fmt.Errorf("无效的平台代码: %v", value)
			}
			return validation.In(
				constant.Shopify,
				constant.Ebay,
				constant.Wish,
				constant.AliExpress,
				constant.Shopee,
				constant.Lazada,
				constant.Walmart,
				constant.CustomPlatform,
				constant.Wayfair,
				constant.TikTok,
			).Error("无效的平台代码：" + code).Validate(code)
		})))),
		validation.Field(&m.StartTime, validation.Required.Error("请输入开始时间"),
			validation.Date(constant.DatetimeFormat).Error("无效的开始时间"),
		),
		validation.Field(&m.EndTime, validation.Required.Error("请输入结束时间"),
			validation.Date(constant.DatetimeFormat).Error("无效的结束时间"),
			validation.By(func(value interface{}) error {
				d1 := m.StartTime
				d2 := value.(string)
				startTime, err := time.Parse(constant.DatetimeFormat, d1)
				if err != nil {
					return err
				}
				endTime, err := time.Parse(constant.DatetimeFormat, d2)
				if err != nil {
					return err
				}
				if startTime.After(endTime) {
					return fmt.Errorf("结束时间不能小于 %s", d1)
				}
				return nil
			}),
		),
	)
}

// All 查询多平台订单列表
// 数据对应多平台管理系统中【订单】>【订单管理】的订单数据，支持查询亚马逊 FBM 订单和多平台订单
func (s multiPlatformOrderService) All(params MultiPlatformOrdersQueryParams) (items []MultiPlatformOrder, nextOffset int, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.SetPagingVars()
	res := struct {
		NormalResponse
		Data struct {
			Current int                  `json:"current"` // 当前页数
			List    []MultiPlatformOrder `json:"list"`    // 详细列表
			Total   int                  `json:"total"`   // 总条数
		} `json:"data"`
	}{}
	resp, err := s.httpClient.R().
		SetBody(params).
		Post("/pb/mp/order/list")
	if err != nil {
		return
	}

	if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
		items = res.Data.List
		nextOffset = params.nextOffset
		isLastPage = len(items) < params.Limit
	}
	return
}
