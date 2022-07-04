package lingxing

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hiscaler/lingxing/entity"
	jsoniter "github.com/json-iterator/go"
)

// 费率
// https://openapidoc.lingxing.com/#/docs/BasicData/Currency

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

func (s basicDataService) Rates(params RatesQueryParams) (items []entity.Rate, nextOffset int, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.SetPagingVars(params.Offset, params.Limit, s.defaultQueryParams.MaxLimit)
	res := struct {
		NormalResponse
		Data []entity.Rate `json:"data"`
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
