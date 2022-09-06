package lingxing

import (
	"encoding/json"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hiscaler/lingxing/constant"
	jsoniter "github.com/json-iterator/go"
	"time"
)

type ProductReport struct {
	SID                         int             `json:"sid"`                            // 店铺 ID
	ID                          int             `json:"id"`                             // ID
	GmtModified                 string          `json:"gmt_modified"`                   // 更新时间
	Price                       float64         `json:"price"`                          // 单价
	Asin                        string          `json:"asin"`                           // ASIN
	SmallImageUrl               string          `json:"small_image_url"`                // 商品图片链接
	ItemName                    string          `json:"item_name"`                      // 标题
	CID                         int             `json:"cid"`                            // 种类 ID
	BID                         int             `json:"bid"`                            // 品牌 ID
	AvailableDays               float64         `json:"avaiable_days"`                  // 可售天数预估
	OrderItems                  int             `json:"order_items"`                    // 订单量
	Volume                      int             `json:"volume"`                         // 销量
	Amount                      float64         `json:"amount"`                         // 销售额
	SessionsBrowser             string          `json:"sessions_browser"`               // Sessions Browser
	SessionsTotal               string          `json:"sessions_total"`                 // Sessions Total
	SessionsMobile              string          `json:"sessions_mobile"`                // Sessions Mobile
	BuyBoxPercentage            float64         `json:"buy_box_percentage"`             // Buy Box
	PageViewsBrowser            float64         `json:"page_views_browser"`             // PV Browser
	PageViewsTotal              float64         `json:"page_views_total"`               // PV Total
	PageViewsMobile             float64         `json:"page_views_mobile"`              // PV Mobile
	Clicks                      int             `json:"clicks"`                         // 点击量
	Impressions                 int             `json:"impressions"`                    // 展示量
	TotalSpend                  float64         `json:"total_spend"`                    // 广告花费
	CTR                         float64         `json:"ctr"`                            // CTR
	AvgCPC                      float64         `json:"avg_cpc"`                        // CPC
	Rank                        int             `json:"rank"`                           // 大类排名
	Reviews                     int             `json:"reviews"`                        // 评论数
	AvgStar                     float64         `json:"avg_star"`                       // 评分
	ConversionRate              float64         `json:"conversion_rate"`                // 转化率
	TotalSpendRate              float64         `json:"total_spend_rate"`               // 总转化率
	AfnFulfillableQuantity      int             `json:"afn_fulfillable_quantity"`       // FBA可售
	ReservedFcTransfers         int             `json:"reserved_fc_transfers"`          // 待调仓
	ReservedFcProcessing        int             `json:"reserved_fc_processing"`         // 调仓中
	ReservedCustomerOrders      int             `json:"reserved_customerorders"`        // 调仓
	AfnInboundShippedQuantity   int             `json:"afn_inbound_shipped_quantity"`   // 入库中
	AfnUnsellableQuantity       int             `json:"afn_unsellable_quantity"`        // 不可售
	AfnInboundReceivingQuantity int             `json:"afn_inbound_receiving_quantity"` // 在途
	AfnInboundWorkingQuantity   int             `json:"afn_inbound_working_quantity"`   // 计划入库
	Acos                        float64         `json:"acos"`                           // ACOS
	Acoas                       float64         `json:"acoas"`                          // ACoAS
	OrderQuantity               float64         `json:"order_quantity"`                 // 广告订单量
	Category                    string          `json:"category"`                       // 类别
	Pid                         int             `json:"pid"`                            // 商品ID
	AdvRate                     float64         `json:"adv_rate"`                       // 广告订单量占比
	SalesAmount                 float64         `json:"sales_amount"`                   // 广告销售额
	AdCvr                       float64         `json:"ad_cvr"`                         // 广告CVR
	Asoas                       float64         `json:"asoas"`                          // ASOAS
	Remark                      json.RawMessage `json:"remark"`                         // asin备注数组，格式[{"date": "", "content": ""}]
	SmallRankList               json.RawMessage `json:"smallRankList"`                  // 小类排名数组，格式[{"smallRankName": "", "rankValue": ""}]
}

type ProductStatisticQueryParams struct {
	Paging
	AsinType  int    `json:"asin_type,omitempty"` // 产品表现维度（0：asin、1：父 asin，不填默认为 0）
	StartDate string `json:"start_date"`          // 报表时间（Y-m-d格式。 eg:2019-07-12，闭区间）
	EndDate   string `json:"end_date"`            // 报表时间（Y-m-d格式。 eg:2019-07-12，开区间）
}

func (m ProductStatisticQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.AsinType, validation.When(!validation.IsEmpty(m.AsinType), validation.In(0, 1))),
		validation.Field(&m.StartDate,
			validation.Required.Error("报表开始时间不能为空"),
			validation.Date(constant.DateFormat).Error("报表开始时间格式有误"),
		),
		validation.Field(&m.EndDate,
			validation.Required.Error("报表结束时间不能为空"),
			validation.Date(constant.DateFormat).Error("报表结束时间格式有误"),
			validation.By(func(value interface{}) error {
				d1 := m.StartDate
				d2 := value.(string)
				startDate, err := time.Parse(constant.DateFormat, d1)
				if err != nil {
					return err
				}
				endDate, err := time.Parse(constant.DateFormat, d2)
				if err != nil {
					return err
				}
				if startDate.After(endDate) {
					return fmt.Errorf("结束时间不能小于 %s", d1)
				}
				return nil
			}),
		),
	)
}

// Products 查询产品表现
// https://openapidoc.lingxing.com/#/docs/Statistics/AsinList
func (s statisticService) Products(params ProductStatisticQueryParams) (items []ProductReport, nextOffset int, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.SetPagingVars()
	res := struct {
		NormalResponse
		Data []ProductReport `json:"data"`
	}{}
	resp, err := s.httpClient.R().
		SetBody(params).
		Post("/data/sales_report/asinList")
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
