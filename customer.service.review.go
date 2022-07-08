package lingxing

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hiscaler/lingxing/constant"
	jsoniter "github.com/json-iterator/go"
)

type customerServiceReviewService service

// Review

type CustomerServiceReview struct {
	Ratings      int      `json:"ratings"`       // 子rating总数
	FiveStar     int      `json:"five_star"`     // 5星review新增数
	FourStar     int      `json:"four_star"`     // 4星review新增数
	ThreeStar    int      `json:"three_star"`    // 3星review新增数
	TwoStar      int      `json:"two_star"`      // 2星review新增数
	OneStar      int      `json:"one_star"`      // 星review新增数
	ReviewNum    int      `json:"review_num"`    // review数
	GoodNum      int      `json:"good_num"`      // review好评数
	NegativeNum  int      `json:"negative_num"`  // review中差评数
	GoodRate     float64  `json:"good_rate"`     // review好评率
	NegativeRate float64  `json:"negative_rate"` // review中差评率
	ModifiedNum  float64  `json:"modified_num"`  // review改评数
	RemoveNum    float64  `json:"remove_num"`    // review删评数
	Asin         string   `json:"asin"`          // asin
	AsinURL      string   `json:"asin_url"`      // asin链接
	ImageURL     string   `json:"image_url"`     // 图片链接
	Title        string   `json:"title"`         // 商品标题
	Country      string   `json:"country"`       // 国家
	Score        float64  `json:"score"`         // 评分
	Mark         float64  `json:"mark"`          // 仅评分数
	SellerName   []string `json:"seller_name"`   // 店铺名称
	LocalInfo    []struct {
		LocalSKU  string `json:"local_sku"`  // SKU
		LocalName string `json:"local_name"` // 品名
	} `json:"local_info"`
	ParentAsin []string `json:"parent_asin"` // 父 SKU
}

type CustomerServiceReviewsQueryParams struct {
	Paging
	SID       string `url:"sid,omitempty"` // 店铺 id
	StartDate string `url:"start_date"`    // 开始时间
	EndDate   string `url:"end_date"`      // 结束时间（最大不超过1年）
}

func (m CustomerServiceReviewsQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.StartDate,
			validation.Required.Error("开始时间不能为空"),
			validation.Date(constant.DateFormat).Error("开始时间格式有误"),
		),
		validation.Field(&m.EndDate,
			validation.Required.Error("结束时间不能为空"),
			validation.Date(constant.DateFormat).Error("结束时间格式有误"),
		),
	)
}

// All 查询 review 列表
// https://openapidoc.lingxing.com/#/docs/Service/reviewLists
func (s customerServiceReviewService) All(params CustomerServiceReviewsQueryParams) (items []CustomerServiceReview, nextOffset int, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.SetPagingVars()
	res := struct {
		NormalResponse
		Data []CustomerServiceReview `json:"data"`
	}{}
	resp, err := s.httpClient.R().
		SetQueryParamsFromValues(toValues(params)).
		Get("/cs/reviewReport/lists")
	if err != nil {
		return
	}

	if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
		items = res.Data
		nextOffset = params.NextOffset
		isLastPage = res.Total < params.Offset
	}
	return
}
