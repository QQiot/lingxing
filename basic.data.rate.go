package lingxing

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	jsoniter "github.com/json-iterator/go"
	"time"
)

// 费率
// https://openapidoc.lingxing.com/#/docs/BasicData/Currency

type Rate struct {
	Date       string    `json:"date"`        // 汇率年月
	Code       string    `json:"code"`        // 币种
	Icon       string    `json:"icon"`        // 币种符号
	Name       string    `json:"name"`        // 币种名
	RateOrg    float64   `json:"rate_org"`    // 官方汇率（数据来源中国银行官方汇率）
	MyRate     float64   `json:"my_rate"`     // 我的汇率（用户自定义汇率，系统首先使用该汇率数据）
	UpdateTime time.Time `json:"update_time"` // 更新时间
}

type RatesQueryParams struct {
	Paging
	Date string `json:"date"` // 汇率月份（格式为：2021-08）
}

func (m RatesQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Date,
			validation.Required.Error("汇率月份不能为空"),
			validation.Date("2006-01").Error("费率月份格式有误，正确的格式为：2006-01"),
		),
	)
}

func (s basicDataService) Rates(params RatesQueryParams) (items []Rate, nextOffset int, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.SetPagingVars()
	res := struct {
		NormalResponse
		Data []Rate `json:"data"`
	}{}
	resp, err := s.httpClient.R().
		SetBody(params).
		Post("/routing/finance/currency/currencyMonth")
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
