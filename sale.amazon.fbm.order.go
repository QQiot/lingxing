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
	WID                   int      `json:"wid"`                     // 发货仓库 ID
	WarehouseName         string   `json:"warehouse_name"`          // 发货仓库名称
	CustomerComment       string   `json:"customer_comment"`        // 客服备注
}

type AmazonFBMOrdersQueryParams struct {
	Paging
	SID         string `json:"sid"`                    // 店铺 ID（多个使用逗号分隔开）
	OrderStatus string `json:"order_status,omitempty"` // 订单状态，用逗号分隔开（2：已发货、3：未付款、4：待审核、5：待发货、6：已取消）
	StartTime   string `json:"start_time,omitempty"`   // 查询时间左闭区间，可精确到时分秒（格式：Y-m-d 或 Y-m-d H:i:s）
	EndTime     string `json:"end_time,omitempty"`     // 查询时间右开区间，可精确到时分秒（格式：Y-m-d 或 Y-m-d H:i:s）
}

func (m AmazonFBMOrdersQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.SID, validation.Required.Error("店铺 ID 不能为空")),
		validation.Field(&m.StartTime, validation.When(m.StartTime != "", validation.Date(constant.DatetimeFormat).Error("查询开始时间格式有误"))),
		validation.Field(&m.EndTime, validation.When(m.EndTime != "", validation.Date(constant.DatetimeFormat).Error("查询结束时间格式有误"))),
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
		nextOffset = params.nextOffset
		isLastPage = res.Total <= params.Offset
	}
	return
}

// 自发货订单详情
// https://openapidoc.lingxing.com/#/docs/Sale/FBMOrderDetail

type FBMOrderDetail struct {
	OrderNumber                  string               `json:"order_number"`                    // 系统单号
	OrderStatus                  string               `json:"order_status"`                    // 订单状态
	OrderFromName                string               `json:"order_from_name"`                 // 订单类型
	PurchaseTime                 string               `json:"purchase_time"`                   // 订购时间
	Platform                     string               `json:"platform"`                        // 平台
	ShopName                     string               `json:"shop_name"`                       // 店铺
	BuyerName                    string               `json:"buyer_name"`                      // 买家姓名（应平台要求，不再返回该数据）
	BuyerEmail                   string               `json:"buyer_email"`                     // 买家邮箱（应平台要求，不再返回该数据）
	BuyerChooseExpress           string               `json:"buyer_choose_express"`            // 客选物流
	TotalShippingPrice           float64              `json:"total_shipping_price"`            // 客付运费
	BuyerMessage                 string               `json:"buyer_message"`                   // 买家留言
	CustomerComment              string               `json:"customer_comment"`                // 客服备注
	Consignee                    string               `json:"consignee"`                       // 收件人（应平台要求，不再返回该数据）
	PostalCode                   string               `json:"postal_code"`                     // 邮编（应平台要求，不再返回该数据）
	Phone                        string               `json:"phone"`                           // 电话（应平台要求，不再返回该数据）
	CountryCode                  string               `json:"country_code"`                    // 国家代码（应平台要求，不再返回该数据）
	StateOrRegion                string               `json:"state_or_region"`                 // 省/州（应平台要求，不再返回该数据）
	City                         string               `json:"city"`                            // 城市（应平台要求，不再返回该数据）
	Address                      string               `json:"address"`                         // 详细地址（应平台要求，不再返回该数据）
	WarehouseName                string               `json:"warehouse_name"`                  // 发货仓库
	WId                          string               `json:"wid"`                             // 发货仓库 ID
	LogisticsTypeName            string               `json:"logistics_type_name"`             // 物流方式
	LogisticsProviderName        string               `json:"logistics_provider_name"`         // 物流商
	LogisticsTypeId              string               `json:"logistics_type_id"`               // 物流方式 ID
	LogisticsProviderId          string               `json:"logistics_provider_id"`           // 物流商 ID
	TrackingNumber               string               `json:"tracking_number"`                 // 跟踪号
	LogisticsPreWeight           float64              `json:"logistics_pre_weight"`            // 估算重量
	LogisticsPreWeightUnit       string               `json:"logistics_pre_weight_unit"`       // 估算重量单位
	PackageLength                float64              `json:"package_length"`                  // 估算尺寸长
	PackageWidth                 float64              `json:"package_width"`                   // 估算尺寸宽
	PackageHeight                float64              `json:"package_height"`                  // 估算尺寸高
	PackageUnit                  string               `json:"package_unit"`                    // 估算尺寸单位
	LogisticsPrePrice            float64              `json:"logistics_pre_price"`             // 预估运费
	PkgRealWeight                float64              `json:"pkg_real_weight"`                 // 包裹实重
	PkgRealWeightUnit            string               `json:"pkg_real_weight_unit"`            // 包裹实重单位
	PkgLength                    float64              `json:"pkg_length"`                      // 包裹尺寸长
	PkgWidth                     float64              `json:"pkg_width"`                       // 包裹尺寸宽
	PkgHeight                    float64              `json:"pkg_height"`                      // 包裹尺寸高
	LogisticsFreight             float64              `json:"logistics_freight"`               // 物流运费
	LogisticsFreightCurrencyCode string               `json:"logistics_freight_currency_code"` // 物流运费币种
	OrderPriceAmount             float64              `json:"order_price_amount"`              // 订单总金额
	GrossProfitAmount            float64              `json:"gross_profit_amount"`             // 订单毛利润
	OrderItem                    []FBMOrderDetailItem `json:"order_item"`                      // 订单项
}

type FBMOrderDetailItem struct {
	PlatformOrderId string   `json:"platform_order_id"` // 平台单号
	MSKU            string   `json:"MSKU"`              // MSKU
	PicURL          string   `json:"pic_url"`           // 图片连接
	SKU             string   `json:"sku"`               // SKU
	ProductName     string   `json:"product_name"`      // 品名
	Quantity        int      `json:"quantity"`          // 数量
	Price           float64  `json:"price"`             // 单价
	CurrencyCode    string   `json:"currency_code"`     // 单价币种
	Customization   string   `json:"customization"`     // 商品备注
	Attachments     []string `json:"attachments"`       // 商品附件
}

func (s fbmOrderService) One(number string) (item FBMOrderDetail, err error) {
	res := struct {
		NormalResponse
		Data FBMOrderDetail `json:"data"`
	}{}
	resp, err := s.httpClient.R().
		SetBody(map[string]string{"order_number": number}).
		Post("/routing/order/Order/getOrderDetail")
	if err != nil {
		return
	}

	if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
		item = res.Data
	}
	return
}
