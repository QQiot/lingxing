package lingxing

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hiscaler/lingxing/constant"
	jsoniter "github.com/json-iterator/go"
	"time"
)

// 仓库
type warehouseService service

// Warehouse 仓库
type Warehouse struct {
	WID  int    `json:"wid"`  // 仓库 ID
	Name string `json:"name"` // 仓库名
	Type int    `json:"type"` // 仓库类型（1；本地、3：海外）
}

type WarehousesQueryParams struct {
	Paging
	Type    int `json:"type,omitempty"`     // 仓库类型（1；本地[默认]、3：海外）
	SubType int `json:"sub_type,omitempty"` // 海外仓子类型（1：无 API 海外仓、2：有 API 海外仓【此参数只有在type=3的时候有效】）
}

func (m WarehousesQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Type, validation.In(1, 3).Error("无效的仓库类型")),
		validation.Field(&m.SubType, validation.When(m.Type == 3, validation.In(1, 2).Error("无效的海外仓子类型"))),
	)
}

// All 查询本地仓库列表
// https://openapidoc.lingxing.com/#/docs/Warehouse/WarehouseLists
func (s warehouseService) All(params WarehousesQueryParams) (items []Warehouse, nextOffset int, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.SetPagingVars()
	res := struct {
		NormalResponse
		Data []Warehouse `json:"data"`
	}{}
	resp, err := s.httpClient.R().
		SetBody(params).
		Post("/data/local_inventory/warehouse")
	if err != nil {
		return
	}

	if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
		items = res.Data
		nextOffset = params.NextOffset
		isLastPage = len(items) < params.Limit
	}
	return
}

// 入库单

// InboundOrderItem 入库单项
type InboundOrderItem struct {
	ProductName     string  `json:"product_name"`      // 品名
	SKU             string  `json:"sku"`               // SKU
	FnSKU           string  `json:"fnsku"`             // FNSKU
	SellerId        string  `json:"seller_id"`         // 店铺id
	Price           float64 `json:"price"`             // 采购单价
	Amount          float64 `json:"amount"`            // 入库成本
	FeeCost         float64 `json:"fee_cost"`          // 费用
	ProductGoodNum  int     `json:"product_good_num"`  // 良品量
	ProductBadNum   int     `json:"product_bad_num"`   // 次品量
	ProductQcNum    int     `json:"product_qc_num"`    // 待检量
	ProductTotal    int     `json:"product_total"`     // 入库量
	ProductAmounts  float64 `json:"product_amounts"`   // 货值
	SingleFee       float64 `json:"single_fee"`        // 单位费用
	SingleStockCost float64 `json:"single_stock_cost"` // 单位入库成本
}

type InboundOrder struct {
	OptRealName     string             `json:"opt_realname"`       // 入库人姓名
	OptTime         string             `json:"opt_time"`           // 	操作时间
	OptUID          int                `json:"opt_uid"`            // 操作人id
	CommitRealName  string             `json:"commit_realname"`    // 提交人名称
	CommitUID       string             `json:"commit_uid"`         // 提交人 ID
	CommitTime      string             `json:"commit_time"`        // 提交时间
	OrderSN         string             `json:"order_sn"`           // 订单号
	Status          int                `json:"status"`             // 入库单状态
	StatusText      string             `json:"status_text"`        // 入库单状态名称
	CreateTime      string             `json:"create_time"`        // 创建时间
	CreateUID       int                `json:"create_uid"`         // 创建人ID
	CreateRealName  string             `json:"create_realname"`    // 创建人名称
	PurchaseOrderSn string             `json:"purchase_order_sn"`  // 采购单号
	RevokeRealName  string             `json:"revoke_realname"`    // 撤销人名称
	RevokeUID       int                `json:"revoke_uid"`         // 撤销人id
	RevokeTime      string             `json:"revoke_time"`        // 撤销时间
	SupplierId      string             `json:"supplier_id"`        // 供应商id
	SupplierName    string             `json:"supplier_name"`      // 供应商名称
	SourceSn        string             `json:"source_sn"`          // 关联单据号
	OrderAmount     float64            `json:"order_amount"`       // 单据入库成本
	CgUID           int                `json:"cg_uid"`             // 采购员 ID
	ReturnPrice     float64            `json:"return_price"`       // 运费
	Currency        string             `json:"currency"`           // 运费币种
	OtherFee        float64            `json:"other_fee"`          // 其他费用
	FeePartType     string             `json:"fee_part_type"`      // 费用分摊方式
	FeePartTypeText string             `json:"fee_part_type_text"` // 费用分摊方式名称
	Type            int                `json:"type"`               // 入库类型
	TypeText        string             `json:"type_text"`          // 入库类型名称
	CgRealName      string             `json:"cg_realname"`        // 采购员姓名
	WID             string             `json:"wid"`                // 仓库 ID
	WarehouseName   string             `json:"ware_house_name"`    // 仓库名称
	Remark          string             `json:"remark"`             // 单据备注
	ItemList        []InboundOrderItem `json:"item_list"`          // 入库项
}

type InboundOrdersQueryParams struct {
	Paging
	WID             string `json:"wid"`               // 系统仓库 ID
	SearchFieldTime string `json:"search_field_time"` // 时间搜索维度（create_time：创建时间、opt_time：入库时间）
	StartDate       string `json:"start_date"`        // 开始日期（Y-m-d，闭区间）
	EndDate         string `json:"end_date"`          // 结束日期（Y-m-d，开区间）
	OrderSn         string `json:"order_sn"`          // 入库单单号，支持多个，分号隔离
	Status          int    `json:"status"`            // 入库单状态（10：待提交、121：待审批、20：待入库、40：已完成、50：已撤销）
	Type            int    `json:"type"`              // 入库类型（1：其他入库、2：采购入库、3：调拨入库、26：退货入库、27：移除入库）
}

func (m InboundOrdersQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.WID, validation.Required.Error("系统仓库 ID 不能为空")),
		validation.Field(&m.SearchFieldTime,
			validation.Required.Error("时间搜索维度不能为空"),
			validation.In("create_time", "opt_time").Error("时间搜索维度有误"),
		),
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
		validation.Field(&m.Status, validation.In(10, 121, 20, 40, 50).Error("无效的入库单状态")),
		validation.Field(&m.Type, validation.In(1, 2, 3, 26, 27).Error("无效的入库类型")),
	)
}

// InboundOrders 获取入库单列表
func (s warehouseService) InboundOrders(params InboundOrdersQueryParams) (items []InboundOrder, nextOffset int, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.SetPagingVars()
	res := struct {
		NormalResponse
		Data []InboundOrder `json:"data"`
	}{}
	resp, err := s.httpClient.R().
		SetBody(params).
		Post("/routing/storage/inbound/getOrders")
	if err != nil {
		return
	}

	if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
		items = res.Data
		nextOffset = params.NextOffset
		isLastPage = len(items) < params.Limit
	}
	return
}
