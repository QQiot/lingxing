package lingxing

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	jsoniter "github.com/json-iterator/go"
)

// Listing

type Listing struct {
	ListingId                   string             `json:"listing_id"`                     // 亚马逊定义的listing的id
	SellerSKU                   string             `json:"seller_sku"`                     // MSKU
	FnSKU                       string             `json:"fnsku"`                          // FNSKU
	ItemName                    string             `json:"item_name"`                      // 品名
	LocalSKU                    string             `json:"local_sku"`                      // 本地SKU
	LocalName                   string             `json:"local_name"`                     // 本地品名
	Price                       float64            `json:"price"`                          // 商品的原价
	Quantity                    int                `json:"quantity"`                       // 商品的数量
	ASIN                        string             `json:"asin"`                           // ASIN
	ParentASIN                  string             `json:"parent_asin"`                    // 父ASIN
	SmallImageURL               string             `json:"small_image_url"`                // 主图URL
	Status                      int                `json:"status"`                         // 状态（1：在售，0：下架）
	IsDelete                    bool               `json:"is_delete"`                      // 是否删除（1：是，0：否）
	AfnFulfillableQuantity      int                `json:"afn_fulfillable_quantity"`       // 可售
	ReservedFcTransfers         int                `json:"reserved_fc_transfers"`          // 待调仓
	ReservedFcProcessing        int                `json:"reserved_fc_processing"`         // 调仓中
	ReservedCustomerOrders      int                `json:"reserved_customerorders"`        // 待发货
	AfnInboundShippedQuantity   int                `json:"afn_inbound_shipped_quantity"`   // 入库
	AfnUnsellableQuantity       int                `json:"afn_unsellable_quantity"`        // 不可售
	AfnInboundWorkingQuantity   int                `json:"afn_inbound_working_quantity"`   // 计划入库
	AfnInboundReceivingQuantity int                `json:"afn_inbound_receiving_quantity"` // 入库中
	CurrencyCode                string             `json:"currency_code"`                  // 币种
	LandedPrice                 float64            `json:"landed_price"`                   // 卖家自己产品的销售价格
	ListingPrice                float64            `json:"listing_price"`                  // listing的显示售价（实际优惠价）
	OpenDate                    string             `json:"open_date"`                      // 商品上架/创建时间
	ListingUpdateDate           string             `json:"listing_update_date"`            // 更新时间
	SellerRank                  int                `json:"seller_rank"`                    // 排名
	SellerCategory              string             `json:"seller_category"`                // 排名所属的类别
	ReviewNum                   int                `json:"review_num"`                     // 评论条数
	LastStar                    float64            `json:"last_star"`                      // 星级评分
	FulfillmentChannelType      string             `json:"fulfillment_channel_type"`       // 配送方式
	PrincipalInfo               []ListingPrincipal `json:"principal_info"`                 // 负责人数据
	Shipping                    float64            `json:"shipping"`                       // 运费
	Points                      float64            `json:"points"`                         // 积分，日本站才有
}

type ListingPrincipal struct {
	UID  string `json:"principal_uid"`  // 负责人用户id
	Name string `json:"principal_name"` // 负责人姓名
}

type ListingsQueryParams struct {
	Paging
	SID    int `json:"sid"`     // 店铺ID
	IsPair int `json:"is_pair"` // 是否配对（1：已配对，2：未配对）
}

func (m ListingsQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.SID, validation.Required.Error("店铺 ID 不能为空")),
		validation.Field(&m.IsPair, validation.In(1, 2).Error("无效的配对状态")),
	)
}

// All 查询listing
// https://openapidoc.lingxing.com/#/docs/Sale/Listing
func (s listingService) All(params ListingsQueryParams) (items []Listing, nextOffset int, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.SetPagingVars()
	res := struct {
		NormalResponse
		Data []Listing `json:"data"`
	}{}
	resp, err := s.httpClient.R().
		SetBody(params).
		Post("/data/mws/listing")
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

// 配对
// https://openapidoc.lingxing.com/#/docs/Sale/Productlink

type ListingPairRequest struct {
	SellerId      string `json:"seller_id,omitempty"`      // 店铺 ID
	MarketplaceId string `json:"marketplace_id,omitempty"` // 市场 ID
	MSKU          string `json:"msku"`                     // MSKU
	SKU           string `json:"sku"`                      // 本地 SKU
	IsSyncPic     bool   `json:"is_sync_pic"`              // 是否同步listing图片（0：否、1：是）
}

func (m ListingPairRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.MSKU, validation.Required.Error("MSKU 不能为空")),
		validation.Field(&m.SKU, validation.Required.Error("本地 SKU 不能为空")),
		validation.Field(&m.IsSyncPic, validation.In(0, 1).Error("是否同步图片标识错误")),
	)
}

func (s listingService) Pair(req ListingPairRequest) (totalCount, successfulCount, failedCount int, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	res := struct {
		NormalResponse
		Data struct {
			Total   int `json:"total"`
			Success int `json:"success"`
			Error   int `json:"error"`
		} `json:"data"`
	}{}
	resp, err := s.httpClient.R().
		SetBody(req).
		Post("/storage/product/link")
	if err != nil {
		return
	}

	if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
		totalCount = res.Data.Success
		successfulCount = res.Data.Success
		failedCount = res.Data.Error
	}
	return
}
