package lingxing

import (
	jsoniter "github.com/json-iterator/go"
	"strings"
)

// 查询亚马逊店铺信息
// https://openapidoc.lingxing.com/#/docs/BasicData/SellerLists

type Seller struct {
	SID             int    `json:"sid"`               // 店铺 ID（领星ERP对企业已授权店铺的唯一标识）
	MID             int    `json:"mid"`               // 站点ID
	Name            string `json:"name"`              // 店铺名
	Country         string `json:"country"`           // 国家
	Region          string `json:"region"`            // 站点简称
	SellerId        string `json:"seller_id"`         // SELLER_ID
	SellerAccountId int    `json:"seller_account_id"` // 销售帐号 ID
	AccountName     string `json:"account_name"`      // 帐号名称
}

type SellersQueryParams struct {
	Name     string `url:"name,omitempty"`      // 店铺名（LIKE）
	SellerId string `url:"seller_id,omitempty"` // Seller ID（EQ）
}

func (m SellersQueryParams) Validate() error {
	return nil
}

// Sellers 查询亚马逊店铺信息
func (s basicDataService) Sellers(params ...SellersQueryParams) (items []Seller, err error) {
	if len(params) > 0 {
		if err = params[0].Validate(); err != nil {
			return
		}
	}

	res := struct {
		NormalResponse
		Data []Seller `json:"data"`
	}{}
	resp, err := s.httpClient.R().Get("/data/seller/lists")
	if err != nil {
		return
	}

	if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
		if len(params) == 0 {
			items = res.Data
		} else {
			name := strings.TrimSpace(params[0].Name)
			sellerId := params[0].SellerId
			if name == "" && sellerId == "" {
				items = res.Data
			} else {
				items = make([]Seller, 0)
				name = strings.ToLower(name)
				for i := range res.Data {
					if (name != "" && !strings.Contains(strings.ToLower(res.Data[i].Name), name)) ||
						sellerId != "" && !strings.EqualFold(res.Data[i].SellerId, sellerId) {
						continue
					}
					items = append(items, res.Data[i])
				}
			}
		}
	}
	return
}
