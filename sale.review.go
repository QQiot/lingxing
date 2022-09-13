package lingxing

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	jsoniter "github.com/json-iterator/go"
)

// https://openapidoc.lingxing.com/#/docs/Sale/Reviews

type reviewService service

type Review struct {
	SID           string   `json:"sid"`             // 店铺 ID
	ASIN          string   `json:"asin"`            // ASIN
	LastStar      float64  `json:"last_star"`       // 星级
	LastTitle     string   `json:"last_title"`      // 标题
	LastContent   string   `json:"last_content"`    // 内容
	Author        string   `json:"author"`          // 评价客户
	AuthorId      string   `json:"author_id"`       // 评价客户 ID
	ReviewDate    string   `json:"review_date"`     // 发表评论日期
	IsVP          bool     `json:"is_vp"`           // 0：没有购买、1：购买评价
	Status        int      `json:"status"`          // 评论处理状态（0：待处理、1：处理中、2：已完成）
	UpdateTime    string   `json:"update_time"`     // 更新时间
	CreateTime    string   `json:"create_time"`     // 创建时间
	IsDelete      bool     `json:"is_delete"`       // 是否删除（0：否、1：是）
	Remark        string   `json:"remark"`          // 评论备注
	OrderIds      string   `json:"order_ids"`       // 订单号
	SmallImageURL string   `json:"small_image_url"` // 图片链接
	History       []string `json:"history"`         // 历史评价
	SellerName    string   `json:"seller_name"`     // 店铺名
	Marketplace   string   `json:"marketplace"`     // 国家
	ASINURL       string   `json:"asin_url"`        // ASIN 链接
	ReviewURl     string   `json:"review_url"`      // 评价链接
}

type ReviewsQueryParams struct {
	Paging
	SID       int    `json:"sid"`        // 店铺 ID
	StartDate string `json:"start_date"` // 评价日期左闭区间（Y-m-d 格式）
	EndDate   string `json:"end_date"`   // 评价日期右开区间（Y-m-d 格式）
}

func (m ReviewsQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.SID, validation.Required.Error("店铺 ID 不能为空")),
	)
}

func (s reviewService) All(params ReviewsQueryParams) (items []Review, nextOffset int, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.SetPagingVars()
	res := struct {
		NormalResponse
		Data []Review `json:"data"`
	}{}
	resp, err := s.httpClient.R().
		SetBody(params).
		Post("/data/mws/reviews")
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
