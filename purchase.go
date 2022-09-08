package lingxing

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hiscaler/lingxing/constant"
	jsoniter "github.com/json-iterator/go"
	"time"
)

// 采购
type purchaseService service

// PurchasePlan 采购计划
type PurchasePlan struct {
	PlanSn               string   `json:"plan_sn"`                 // 采购计划编号
	StatusText           string   `json:"status_text"`             // 状态
	Status               int      `json:"status"`                  // 计划状态码
	CreatorRealName      string   `json:"creator_real_name"`       // 创建人名称
	CreatorUid           int      `json:"creator_uid"`             // 创建人ID
	CreateTime           string   `json:"create_time"`             // 创建时间 Y-m-d H:i:s
	File                 []string `json:"file"`                    // 附件
	PlanRemark           string   `json:"plan_remark"`             // 备注
	PicUrl               string   `json:"pic_url"`                 // 产品图片
	SpuName              string   `json:"spu_name"`                // 款名
	Spu                  string   `json:"spu"`                     // SPU
	ProductName          string   `json:"product_name"`            // 品名
	ProductId            int      `json:"product_id"`              // 商品ID
	SKU                  string   `json:"sku"`                     // SKU
	Attribute            []string `json:"attribute"`               // 属性
	SID                  int      `json:"sid"`                     // 店铺ID
	SellerName           string   `json:"seller_name"`             // 店铺名称
	Marketplace          string   `json:"marketplace"`             // 国家
	FNSKU                string   `json:"fnsku"`                   // FNSKU
	MSKU                 []string `json:"msku"`                    // MSKU
	SupplierId           string   `json:"supplier_id"`             // 供应商ID
	SupplierName         string   `json:"supplier_name"`           // 供应商名称
	WID                  int      `json:"wid"`                     // 仓库ID
	WarehouseName        string   `json:"warehouse_name"`          // 仓库名称
	PurchaserId          int      `json:"purchaser_id"`            // 采购方ID
	PurchaserName        string   `json:"purchaser_name"`          // 采购方名称
	CgBoxPcs             int      `json:"cg_box_pcs"`              // 单箱数量
	QuantityPlan         int      `json:"quantity_plan"`           // 计划采购量
	ExpectArriveTime     string   `json:"expect_arrive_time"`      // 期望到货时间 Y-m-d
	CgUid                int      `json:"cg_uid"`                  // 采购员ID
	CgOptUsername        string   `json:"cg_opt_username"`         // 采购员名称
	Remark               string   `json:"remark"`                  // 产品备注
	IsCombo              bool     `json:"is_combo"`                // 是否为组合商品。1：是，0：否
	IsAux                bool     `json:"is_aux"`                  // 是否为辅料。1：是，0：否
	IsRelatedProcessPlan bool     `json:"is_related_process_plan"` // 是否关联了加工计划。1：是，0：否
}

type PurchasePlansQueryParams struct {
	Paging
	SearchFieldTime      string   `json:"search_field_time,omitempty"`       // 时间搜索维度：creator_time：创建时间；expect_arrive_time：预计到货时间
	StartDate            string   `json:"start_date"`                        // 开始日期，Y-m-d，闭区间
	EndDate              string   `json:"end_date"`                          // 结束日期，Y-m-d，闭区间
	PlanSNs              []string `json:"plan_sns,omitempty"`                // 采购计划编号
	IsCombo              bool     `json:"is_combo,omitempty"`                // 是否为组合商品（0：否、1：是）
	IsRelatedProcessPlan bool     `json:"is_related_process_plan,omitempty"` // 是否关联加工计划（0：否、1：是）
	Status               []int    `json:"status,omitempty"`                  // 状态（待采购=2 ，已处理=-2，已驳回=122，已作废=-3,124）
	SIDs                 []int    `json:"sids,omitempty"`                    // 店铺
}

func (m PurchasePlansQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.StartDate,
			validation.Required.Error("开始时间不能为空"),
			validation.Date(constant.DateFormat).Error("开始时间格式有误"),
		),
		validation.Field(&m.EndDate,
			validation.Required.Error("结束时间不能为空"),
			validation.Date(constant.DateFormat).Error("结束时间格式有误"),
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

// Plans 查询采购计划列表
// https://openapidoc.lingxing.com/#/docs/Purchase/getPurchasePlans?id=%e6%9f%a5%e8%af%a2%e9%87%87%e8%b4%ad%e8%ae%a1%e5%88%92%e5%88%97%e8%a1%a8
func (s purchaseService) Plans(params PurchasePlansQueryParams) (items []PurchasePlan, nextOffset int, isLastPage bool, err error) {
	params.SetPagingVars()
	if err = params.Validate(); err != nil {
		return
	}

	params.SetPagingVars()
	res := struct {
		NormalResponse
		Data []PurchasePlan `json:"data"`
	}{}
	resp, err := s.httpClient.R().
		SetBody(params).
		Post("/routing/data/local_inventory/getPurchasePlans")
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
