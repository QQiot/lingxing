package lingxing

import (
	"context"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hiscaler/lingxing/constant"
	jsoniter "github.com/json-iterator/go"
)

// https://openapidoc.lingxing.com/#/docs/MultiPlatform/StoreInfo

type multiPlatformSellerService service

// MultiPlatformSeller 多平台店铺信息
type MultiPlatformSeller struct {
	Currency     string `json:"currency"`      // 店铺币种
	PlatformCode string `json:"platform_code"` // 平台 code
	PlatformName string `json:"platform_name"` // 平台名称
	StoreId      int    `json:"store_id"`      // 店铺 ID
	StoreName    string `json:"store_name"`    // 店铺名称
}

type MultiPlatformSellersQueryParams struct {
	Paging
	PlatformCode []string `json:"platform_code,omitempty"` // 平台代码
}

func (m MultiPlatformSellersQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
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
	)
}

// All 查询多平台店铺信息
func (s multiPlatformSellerService) All(params MultiPlatformSellersQueryParams) (items []MultiPlatformSeller, nextOffset int, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.SetPagingVars()
	res := struct {
		NormalResponse
		Data struct {
			Current int                   `json:"current"` // 当前页数
			List    []MultiPlatformSeller `json:"list"`    // 详细列表
			Total   int                   `json:"total"`   // 总条数
		} `json:"data"`
	}{}
	resp, err := s.httpClient.R().
		SetBody(params).
		Post("/pb/mp/shop/getSellerList")
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
