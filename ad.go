package lingxing

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hiscaler/lingxing/constant"
	jsoniter "github.com/json-iterator/go"
	"time"
)

type adService service

type AdGroup struct {
	CampaignId    string  `json:"campaign_id"`    // 广告活动ID
	AdGroupId     string  `json:"ad_group_id"`    // 广告组ID
	AdGroupName   string  `json:"ad_group_name"`  // 广告组名称
	State         string  `json:"state"`          // 广告组状态
	ServingStatus string  `json:"serving_status"` // 广告组的服务状态
	DefaultBid    float64 `json:"default_bid"`    // 竞价
	TargetingMode int     `json:"targeting_mode"` // 广告组投放模式（1：关键词投放、2：商品投放、T00020：SD 的商品投放、T00030：SD 的受众投放）
	TargetingType string  `json:"targeting_type"` // 广告活动匹配类型，如自动或手动
	ProductNum    int     `json:"product_num"`    // 广告数
	Impressions   int     `json:"impressions"`    // 展示
	Clicks        int     `json:"clicks"`         // 点击
	Cost          float64 `json:"cost"`           // 花费
	OrderNum      int     `json:"order_num"`      // 订单量
	SalesAmount   float64 `json:"sales_amount"`   // 销售额
	CurrencyCode  string  `json:"currency_code"`  // 币种
	CampaignName  string  `json:"campaign_name"`  // 广告活动名称
	CTR           float64 `json:"ctr"`            // CTR
	CPC           float64 `json:"cpc"`            // CPC
	CVR           float64 `json:"cvr"`            // CVR
	CPA           float64 `json:"cpa"`            // CPA
	ACOS          float64 `json:"acos"`           // ACOS
	ROAS          float64 `json:"roas"`           // ROAS（SP广告暂无）
}

type AdGroupsQueryParams struct {
	Paging
	SID       int    `json:"sid"`        // 店铺 ID
	StartDate string `json:"start_date"` // 广告时间左闭区间，Y-m-d格式
	EndDate   string `json:"end_date"`   // 广告时间右开区间，Y-m-d格式
	Type      int    `json:"type"`       //	广告类型（1：SP、3：SD）
}

func (m AdGroupsQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.SID, validation.Required.Error("店铺 ID 不能为空")),
		validation.Field(&m.StartDate,
			validation.Required.Error("广告开始时间不能为空"),
			validation.Date(constant.DateFormat).Error("广告开始时间格式有误"),
		),
		validation.Field(&m.EndDate,
			validation.Required.Error("广告结束时间不能为空"),
			validation.Date(constant.DateFormat).Error("广告结束时间格式有误"),
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
		validation.Field(&m.Type,
			validation.Required.Error("广告类型不能为空"),
			validation.In(1, 3).Error("无效的广告类型"),
		),
	)
}

// Groups 查询广告管理-广告组
// https://openapidoc.lingxing.com/#/docs/Advertisement/AdManageGroups
func (s adService) Groups(params AdGroupsQueryParams) (items []AdGroup, nextOffset int, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.SetPagingVars()
	res := struct {
		NormalResponse
		Data []AdGroup `json:"data"`
	}{}
	resp, err := s.httpClient.R().
		SetBody(params).
		Post("/data/ads/adGroups")
	if err != nil {
		return
	}

	if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
		items = res.Data
		nextOffset = params.NextOffset
		isLastPage = res.Total <= nextOffset
	}
	return
}
