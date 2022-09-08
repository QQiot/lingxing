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
	CreatorUID           int      `json:"creator_uid"`             // 创建人 ID
	CreateTime           string   `json:"create_time"`             // 创建时间 Y-m-d H:i:s
	File                 []string `json:"file"`                    // 附件
	PlanRemark           string   `json:"plan_remark"`             // 备注
	PicUrl               string   `json:"pic_url"`                 // 产品图片
	SpuName              string   `json:"spu_name"`                // 款名
	Spu                  string   `json:"spu"`                     // SPU
	ProductName          string   `json:"product_name"`            // 品名
	ProductId            int      `json:"product_id"`              // 商品 ID
	SKU                  string   `json:"sku"`                     // SKU
	Attribute            []string `json:"attribute"`               // 属性
	SID                  int      `json:"sid"`                     // 店铺 ID
	SellerName           string   `json:"seller_name"`             // 店铺名称
	Marketplace          string   `json:"marketplace"`             // 国家
	FNSKU                string   `json:"fnsku"`                   // FNSKU
	MSKU                 []string `json:"msku"`                    // MSKU
	SupplierId           string   `json:"supplier_id"`             // 供应商 ID
	SupplierName         string   `json:"supplier_name"`           // 供应商名称
	WID                  int      `json:"wid"`                     // 仓库 ID
	WarehouseName        string   `json:"warehouse_name"`          // 仓库名称
	PurchaserId          int      `json:"purchaser_id"`            // 采购方 ID
	PurchaserName        string   `json:"purchaser_name"`          // 采购方名称
	CgBoxPcs             int      `json:"cg_box_pcs"`              // 单箱数量
	QuantityPlan         int      `json:"quantity_plan"`           // 计划采购量
	ExpectArriveTime     string   `json:"expect_arrive_time"`      // 期望到货时间（Y-m-d）
	CgUID                int      `json:"cg_uid"`                  // 采购员 ID
	CgOptUsername        string   `json:"cg_opt_username"`         // 采购员名称
	Remark               string   `json:"remark"`                  // 产品备注
	IsCombo              bool     `json:"is_combo"`                // 是否为组合商品（0：否、1：是）
	IsAux                bool     `json:"is_aux"`                  // 是否为辅料（0：否、1：是）
	IsRelatedProcessPlan bool     `json:"is_related_process_plan"` // 是否关联了加工计划（0：否、1：是）
}

type PurchasePlansQueryParams struct {
	Paging
	SearchFieldTime      string   `json:"search_field_time,omitempty"`       // 时间搜索维度（creator_time：创建时间、expect_arrive_time：预计到货时间）
	StartDate            string   `json:"start_date"`                        // 开始日期（Y-m-d，闭区间）
	EndDate              string   `json:"end_date"`                          // 结束日期（Y-m-d，开区间）
	PlanSNs              []string `json:"plan_sns,omitempty"`                // 采购计划编号
	IsCombo              bool     `json:"is_combo,omitempty"`                // 是否为组合商品（0：否、1：是）
	IsRelatedProcessPlan bool     `json:"is_related_process_plan,omitempty"` // 是否关联加工计划（0：否、1：是）
	Status               []int    `json:"status,omitempty"`                  // 状态（待采购=2 ，已处理=-2，已驳回=122，已作废=-3,124）
	SIDs                 []int    `json:"sids,omitempty"`                    // 店铺
}

func (m PurchasePlansQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.SearchFieldTime, validation.When(m.SearchFieldTime != "", validation.In("creator_time", "expect_arrive_time").Error("时间搜索维度有误"))),
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

// 采购单处理

// PurchaseOrderItem 采购单子项
type PurchaseOrderItem struct {
	ID                string   `json:"id"`                  // 子项 ID
	PlanSn            string   `json:"plan_sn"`             // 采购计划号
	ProductId         int      `json:"product_id"`          // 商品 ID
	ProductName       string   `json:"product_name"`        // 品名
	SKU               string   `json:"sku"`                 // SKU
	FNSKU             string   `json:"fnsku"`               // FNSKU
	SID               string   `json:"sid"`                 // 店铺 ID
	Price             float64  `json:"price"`               // 含税单价
	Amount            float64  `json:"amount"`              // 价税合计
	QuantityPlan      int      `json:"quantity_plan"`       // 计划采购量
	QuantityReal      int      `json:"quantity_real"`       // 实际采购量
	QuantityEntry     int      `json:"quantity_entry"`      // 到货入库量
	QuantityReceive   int      `json:"quantity_receive"`    // 待到货量
	QuantityQc        int      `json:"quantity_qc"`         // 质检量
	QuantityQcPrepare int      `json:"quantity_qc_prepare"` // 待质检量
	ExpectArriveTime  string   `json:"expect_arrive_time"`  // 期待到货时间
	Remark            string   `json:"remark"`              // 备注
	CasesNum          int      `json:"cases_num"`           // 箱数
	QuantityPerCase   int      `json:"quantity_per_case"`   // 单箱数量
	IsDelete          bool     `json:"is_delete"`           // 是否删除（0：否、1：是）
	QuantityReturn    int      `json:"quantity_return"`     // 退货数
	MSKU              []string `json:"msku"`                // MSKU
	Attribute         []string `json:"attribute"`           // 属性
	TaxRate           string   `json:"tax_rate"`            // 税率
	SPU               string   `json:"spu"`                 // spu
	SPUName           string   `json:"spu_name"`            // 款名
}

// PurchaseOrderLogisticsInformation 物流信息
type PurchaseOrderLogisticsInformation struct {
	LogisticsCompany string `json:"logistics_company"`  // 物流公司
	LogisticsOrderNo string `json:"logistics_order_no"` // 物流单号
	PolId            string `json:"pol_id"`             // 物流信息记录ID
	PurchaseOrderId  string `json:"purchase_order_id"`  // 采购订单唯一ID
	PurchaseOrderSn  string `json:"purchase_order_sn"`  // 采购订单号（order_sn）
}

// PurchaseOrder 采购单
type PurchaseOrder struct {
	OrderSn               string                              `json:"order_sn"`               // 采购单号
	SupplierId            int                                 `json:"supplier_id"`            // 供应商 ID
	SupplierName          string                              `json:"supplier_name"`          // 供应商
	OptUID                int                                 `json:"opt_uid"`                // 操作员 UID
	CreateTime            string                              `json:"create_time"`            // 创建时间
	OrderTime             string                              `json:"order_time"`             // 下单时间
	PurchaseCurrency      string                              `json:"purchase_currency"`      // 采购币种
	ShippingCurrency      string                              `json:"shipping_currency"`      // 运费币种
	PurchaseRate          float64                             `json:"purchase_rate"`          // 采购汇率
	StatusShipped         int                                 `json:"status_shipped"`         // 到货状态（1：未到货、2：部分到货、3：全部到货）
	QuantityTotal         int                                 `json:"quantity_total"`         // 采购总量
	Payment               float64                             `json:"payment"`                // 应付货款（手工）
	AuditorUID            int                                 `json:"auditor_uid"`            // 审核人员 ID
	AuditorTime           string                              `json:"auditor_time"`           // 审核时间
	LastUID               int                                 `json:"last_uid"`               // 最后操作人员 ID
	LastTime              string                              `json:"last_time"`              // 最后操作时间
	Reason                string                              `json:"reason"`                 // 作废原因
	WID                   int                                 `json:"wid"`                    // 仓库 ID
	IsTax                 bool                                `json:"is_tax"`                 // 是否含税（0：否、1：是）
	Status                int                                 `json:"status"`                 // 状态（-1：作废、0：待审核 - 草稿、1：待下单 - 已审核、2：待签收(待到货) - 已下单、9：完成、121：(审批流)待审核、122：(审批流)驳回、124：(审批流)作废）
	WareHouseBakName      string                              `json:"ware_house_bak_name"`    // 仓库名(备份)
	StatusText            string                              `json:"status_text"`            // 状态文本
	PayStatusText         string                              `json:"pay_status_text"`        // 支付状态文本
	AuditorRealName       string                              `json:"auditor_realname"`       // 审核人姓名
	OptRealNAme           string                              `json:"opt_realname"`           // 操作人姓名
	StatusShippedText     string                              `json:"status_shipped_text"`    // 到货状态文本
	LastRealName          string                              `json:"last_realname"`          // 最后操作人姓名
	ShippingPrice         float64                             `json:"shipping_price"`         // 运费
	AmountTotal           float64                             `json:"amount_total"`           // 货物总价
	PayStatus             int                                 `json:"pay_status"`             // 付款状态（0：未申请、1：已申请、2：部分付款、3：已付款）
	Remark                string                              `json:"remark"`                 // 备注
	OtherFee              float64                             `json:"other_fee"`              // 其他费用
	OtherCurrency         string                              `json:"other_currency"`         // 其他费用币种
	FeePartType           int                                 `json:"fee_part_type"`          // 费用分摊方式（0：不分摊、1：按金额、2：按数量）
	TotalPrice            float64                             `json:"total_price"`            // 总金额
	ICON                  string                              `json:"icon"`                   // 采购币种符号
	WareHouseName         string                              `json:"ware_house_name"`        // 仓库名
	QuantityEntry         int                                 `json:"quantity_entry"`         // 入库量
	QuantityReal          int                                 `json:"quantity_real"`          // 实际采购量
	QuantityReceive       int                                 `json:"quantity_receive"`       // 待到货量
	UpdateTime            string                              `json:"update_time"`            // 采购单更新时间
	ItemList              []PurchaseOrderItem                 `json:"item_list"`              // 采购单子项
	LogisticsInfo         []PurchaseOrderLogisticsInformation `json:"logistics_info"`         // 物流信息
	PurchaserId           int                                 `json:"purchaser_id"`           // 采购方 ID
	ContactPerson         string                              `json:"contact_person"`         // 联系人
	ContactNumber         string                              `json:"contact_number"`         // 联系方式
	SettlementMethod      int                                 `json:"settlement_method"`      // 结算方式
	SettlementDescription string                              `json:"settlement_description"` // 结算描述
	PaymentMethod         int                                 `json:"payment_method"`         // 支付方式
}

type PurchaseOrdersQueryParams struct {
	Paging
	SearchFieldTime string `json:"search_field_time,omitempty"` // 时间搜索维度（create_time：创建时间、expect_arrive_time：预计到货时间）
	StartDate       string `json:"start_date"`                  // 开始日期（Y-m-d，闭区间）
	EndDate         string `json:"end_date"`                    // 结束日期（Y-m-d，开区间）
}

func (m PurchaseOrdersQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.SearchFieldTime, validation.When(m.SearchFieldTime != "", validation.In("create_time", "expect_arrive_time").Error("时间搜索维度有误"))),
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

func (s purchaseService) Orders(params PurchaseOrdersQueryParams) (items []PurchaseOrder, nextOffset int, isLastPage bool, err error) {
	params.SetPagingVars()
	if err = params.Validate(); err != nil {
		return
	}

	params.SetPagingVars()
	res := struct {
		NormalResponse
		Data []PurchaseOrder `json:"data"`
	}{}
	resp, err := s.httpClient.R().
		SetBody(params).
		Post("/routing/data/local_inventory/purchaseOrderList")
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
