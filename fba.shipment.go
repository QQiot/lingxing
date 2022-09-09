package lingxing

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hiscaler/lingxing/constant"
	jsoniter "github.com/json-iterator/go"
	"time"
)

type fbaShipmentService service

// 查询 FBA 发货单列表

type FBAShipment struct {
	ID                   int    `json:"id"`                     // 发货单id
	ShipmentSn           string `json:"shipment_sn"`            // 发货单号
	Status               int    `json:"status"`                 // 发货单状态，-1 : 待配货 0：待发货，1：已发货，2：已完成，3：已作废
	ShipmentTime         string `json:"shipment_time"`          // 发货时间
	WName                string `json:"wname"`                  // 仓库名称
	CreateUser           string `json:"create_user"`            // 创建用户
	LogisticsChannelName string `json:"logistics_channel_name"` // 物流方式
	ExpectedArrivalDate  string `json:"expected_arrival_date"`  // 到货时间
	EtaDate              string `json:"eta_date"`               // 预计到港时间
	DeliveryDate         string `json:"delivery_date"`          // 实际妥投时间
	CreateTime           string `json:"create_time"`            // 创建时间
	IsPick               bool   `json:"is_pick"`                // 拣货状态 0 未拣货 1已拣货
	IsPrint              bool   `json:"is_print"`               // 是否打印
	PickTime             string `json:"pick_time"`              // 拣货时间
	PrintNum             int    `json:"print_num"`              // 打印次数
	HeadFeeType          int    `json:"head_fee_type"`          // 头程费分配方式，0：按计费重；1：按实重；2：按体积重；3：按SKU数量；4：自定义；5：按箱子体积
	FileId               string `json:"file_id"`                // 附件文件
	GmtModified          string `json:"gmt_modified"`           // 更新时间
	Remark               string `json:"remark"`                 // 备注
	Wid                  int    `json:"wid"`                    // 仓库ID
	IsReturnStock        bool   `json:"is_return_stock"`        // 是否恢复库存
	Logistics            []struct {
		ReplaceTrackingNumber string `json:"replace_tracking_number"` // 跟踪单号
		TrackingNumber        string `json:"tracking_number"`         // 物流商号
	} `json:"logistics"` // 	物流列表
	RelateList []struct {
		Mid                            int      `json:"mid"`                               // 国家id
		DestinationFulfillmentCenterId string   `json:"destination_fulfillment_center_id"` // 物流中心编码
		QuantityShipped                int      `json:"quantity_shipped"`                  // 申报量
		Id                             int      `json:"id"`                                // 明细id
		WName                          string   `json:"wname"`                             // 仓库名称
		ShipmentSn                     string   `json:"shipment_sn"`                       // 发货单号
		ShipmentId                     string   `json:"shipment_id"`                       // 货件id
		Wid                            int      `json:"wid"`                               // 仓库id
		Pid                            int      `json:"pid"`                               // 货件明细ID
		SName                          string   `json:"sname"`                             // 店铺名称
		ProductName                    string   `json:"product_name"`                      // 产品名称
		Num                            int      `json:"num"`                               // 发货数量
		PicURL                         string   `json:"pic_url"`                           // 图片url
		PackingType                    int      `json:"packing_type"`                      // 混装类型 2原装 1原装
		FulfillmentNetworkSku          string   `json:"fulfillment_network_sku"`           // listing的fnsku
		Sku                            string   `json:"sku"`                               // sku
		FnSku                          string   `json:"fnsku"`                             // 仓库fnsku
		MSku                           string   `json:"msku"`                              // seller_sku
		Nation                         string   `json:"nation"`                            // 国家名称
		ApplyNum                       int      `json:"apply_num"`                         // 关联货件量
		ProductId                      int      `json:"product_id"`                        // 商品id
		Remark                         string   `json:"remark"`                            // 备注
		Status                         int      `json:"stauts"`                            // 状态
		SId                            int      `json:"sid"`                               // 店铺id
		IsCombo                        bool     `json:"is_combo"`                          // 是否组合商品
		CreateByMws                    int      `json:"create_by_mws"`                     // 创建发货单的途径
		WhbCodeList                    []string `json:"whb_code_list"`                     // 仓位编码列表
		PackingTypeName                string   `json:"packing_type_name"`                 // 包装名称
		ProductValidNum                int      `json:"product_valid_num"`                 // 可用量
		ProductQCNum                   int      `json:"product_qc_num"`                    // 待检量
		DiffNum                        int      `json:"diff_num"`                          // 差额
	} `json:"relate_list"` // 	关联货件列表
	NotRelateList                  []string `json:"not_relate_list"`                   // 未关联货件列表
	DestinationFulfillmentCenterId string   `json:"destination_fulfillment_center_id"` // 物流中心编码
	StatusName                     string   `json:"status_name"`                       // 状态名称
	HeadFeeTypeName                string   `json:"head_fee_type_name"`                // 头程分摊名称
	FileList                       []string `json:"fileList"`                          // 文件列表
}

type FBAShipmentsQueryParams struct {
	Paging
	SearchValue   string   `json:"search_value,omitempty"`   // 搜索的值
	SearchField   string   `json:"search_field,omitempty"`   // 搜索字段，shipment_sn:发货单号，sku:skushipment_id：货件单号
	SIds          []string `json:"sids,omitempty"`           // 店铺id
	MIds          []string `json:"mids,omitempty"`           // 国家id
	WId           []string `json:"wid,omitempty"`            // 仓库id
	LogisticsType []string `json:"logistics_type,omitempty"` // 物流方式id
	Status        int      `json:"status,omitempty"`         // 发货单状态，-1 : 待配货 0：待发货，1：已发货，3：已作废
	PrintStatus   int      `json:"print_status,omitempty"`   // 打印状态： 0未打印 ，1 已打印
	PickStatus    int      `json:"pick_status,omitempty"`    // 拣货状态 ：0 未拣货， 1已拣货
	TimeType      int      `json:"time_type,omitempty"`      // 按时间查询时必传。时间类型 2创建时间 1到货时间 0发货时间
	StartDate     string   `json:"start_date,omitempty"`     // 开始日期
	EndDate       string   `json:"end_date,omitempty"`       // 结束日期
}

func (m FBAShipmentsQueryParams) Validate() error {
	return nil
}

// All 查询 FBA 发货单
func (s fbaShipmentService) All(params FBAShipmentsQueryParams) (items []FBAShipment, nextOffset int, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.SetPagingVars()
	res := struct {
		NormalResponse
		Data []FBAShipment `json:"data"`
	}{}
	resp, err := s.httpClient.R().
		SetBody(params).
		Post("/storage/shipment/getInboundShipmentList")
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

// 查询 FBA 发货单详情
// https://openapidoc.lingxing.com/#/docs/FBA/getInboundShipmentListMwsDetail

type FBAShipmentDetail struct {
	ID           int    `json:"id"`            // 发货单ID
	ZId          int    `json:"zid"`           // ZID
	TrackingId   int    `json:"tracking_id"`   // 物流追踪(运单)ID
	ShipmentSn   string `json:"shipment_sn"`   // 发货单号
	Status       int    `json:"status"`        // 发货单状态，1 : 待配货 0：待发货，1：已发货，2：已完成，3：已作废
	ShipmentTime string `json:"shipment_time"` // 发货时间
	WId          int    `json:"wid"`           // 仓库ID
	GmtModified  string `json:"gmt_modified"`  // 修改时间
	GmtCreate    string `json:"gmt_create"`    // 创建时间
	Remark       string `json:"remark"`        // 备注
	// todo 待完善
}

func (s fbaShipmentService) One(shipmentSN string) (item FBAShipmentDetail, err error) {
	res := struct {
		NormalResponse
		Data FBAShipmentDetail `json:"data"`
	}{}
	_, err = s.httpClient.R().
		SetBody(map[string]string{"shipment_sn": shipmentSN}).
		SetResult(&res).
		Post("/routing/storage/shipment/getInboundShipmentListMwsDetail")
	if err != nil {
		return
	}
	item = res.Data
	return
}

// FBA 发货计划

// FBAShipmentPlan FBA 发货计划
type FBAShipmentPlan struct {
	IspgId     int    `json:"ispg_id"`     // 发货计划组 ID
	CreateTime string `json:"create_time"` // 创建时间
	Seq        string `json:"seq"`         // 批次号
	Remark     string `json:"remark"`      // 备注
	CreateUser string `json:"create_user"` // 创建用户
	List       []struct {
		IspgId               int    `json:"ispg_id"`                // 发货计划组父 ID
		IspId                int    `json:"isp_id"`                 // 发货计划 ID
		LogisticsChannelId   int    `json:"logistics_channel_id"`   // 物流 ID
		Fnsku                string `json:"fnsku"`                  // FNSKU
		Msku                 string `json:"msku"`                   // MSKU
		WID                  int    `json:"wid"`                    // 仓库 ID
		WarehouseName        string `json:"wname"`                  // 仓库名称
		SID                  int    `json:"sid"`                    // 店铺 ID
		CreateTime           string `json:"create_time"`            // 创建时间
		Status               int    `json:"status"`                 // 状态（-5：已驳回、0：待审核、5：待处理、10：已处理）
		PackageType          int    `json:"package_type"`           // 包装类型 2原装 1混装
		ShipmentTime         string `json:"shipment_time"`          // 计划发货时间
		ShipmentPlanQuantity int    `json:"shipment_plan_quantity"` // 计划发货量
		Seq                  string `json:"seq"`                    // 批次号
		LogisticsName        string `json:"logistics_name"`         // 物流名称
		QuantityInCase       int    `json:"quantity_in_case"`       // 单箱数量
		BoxNum               int    `json:"box_num"`                // 箱数
		IsRelateMws          int    `json:"is_relate_mws"`          // 是否关联货件
		IsRelateList         int    `json:"is_relate_list"`         // 是否关联发货单
		Remark               string `json:"remark"`                 // 备注
		PrintNum             int    `json:"print_num"`              // 打印次数
		CreateUser           string `json:"create_user"`            // 创建用户
		SmallImageURL        string `json:"small_image_url"`        // 商品图片
		OrderSN              string `json:"order_sn"`               // 计划发货单号
		ProductName          string `json:"product_name"`           // 产品名称
		ProductId            int    `json:"product_id"`             // 产品 ID
		SKU                  string `json:"sku"`                    // SKU
		PicURL               string `json:"pic_url"`                // 商品图片
		IsCombo              bool   `json:"is_combo"`               // 是否组合商品
		StorageList          []struct {
			ProductId       int `json:"product_id"`        // 商品 ID
			ProductValidNum int `json:"product_valid_num"` // 库存可用量
			ProductQcNum    int `json:"product_qc_num"`    // 待检量
			QuantityReceive int `json:"quantity_receive"`  // 待收货量
		} `json:"storage_list"` // 库存列表
		MwsRelate []struct {
			IsimId               int    `json:"isim_id"`                // 关联货件 ID
			IsilId               int    `json:"isil_id"`                // 关联发货单 ID
			IsilmId              int    `json:"isilm_id"`               // 关联发货单明细 ID
			ShipmentPlanSn       string `json:"shipment_plan_sn"`       // 关联发货计划单号
			ShipmentMwsSn        string `json:"shipment_mws_sn"`        // 关联货件单号
			ShipmentListSn       string `json:"shipment_list_sn"`       // 关联发货单单号
			ShipmentPlanQuantity int    `json:"shipment_plan_quantity"` // 关联发货计划数量
			ShipmentMwsQuantity  int    `json:"shipment_mws_quantity"`  // 关联货件数量
			ShipmentListQuantity int    `json:"shipment_list_quantity"` // 关联发货单数量
		} `json:"mws_relate"` // 发货单关联
		StatusName      string `json:"status_name"`       // 状态名称
		PackingTypeName string `json:"packing_type_name"` // 包装类型名称
		DiffNum         int    `json:"diff_num"`          // 差额
		SName           string `json:"sname"`             // 店铺
		Nation          string `json:"nation"`            // 国家
	} `json:"list"` // 子项目列表
}

type FBAShipmentPlansQueryParams struct {
	Paging
	SIDs            string `json:"sids"`                 // 店铺 ids，多个之间以逗号分隔
	WID             string `json:"wid"`                  // 仓库 ID
	PackageType     int    `json:"package_type"`         // 包装类型（1：混装、2：原装）
	SearchFieldTime string `json:"search_field_time"`    // 查找时间字段（gmt_create?）文档不全
	SearchField     string `json:"search_field"`         // 查找字段（order_sn：发货计划单号）文档不全
	SearchValue     string `json:"search_value"`         // 查找值
	Status          string `json:"status"`               // 状态
	MIDs            string `json:"mids"`                 // 国家 ID
	StartDate       string `json:"start_date,omitempty"` // 开始日期
	EndDate         string `json:"end_date,omitempty"`   // 结束日期
}

func (m FBAShipmentPlansQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.SIDs, validation.Required.Error("店铺 ID 不能为空")),
		validation.Field(&m.WID, validation.Required.Error("仓库 ID 不能为空")),
		validation.Field(&m.PackageType, validation.In(1, 2).Error("无效的包装类型")),
		validation.Field(&m.SearchFieldTime,
			validation.Required.Error("查找时间字段不能为空"),
			validation.In("gmt_create").Error("查找时间字段有误"),
		),
		validation.Field(&m.SearchField,
			validation.Required.Error("查找字段不能为空"),
			validation.In("order_sn").Error("查找字段有误"),
		),
		validation.Field(&m.SearchValue, validation.Required.Error("查找值不能为空")),
		validation.Field(&m.Status, validation.Required.Error("状态不能为空")),
		validation.Field(&m.MIDs, validation.Required.Error("国家 ID 不能为空")),
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

// Plans 查询FBA发货计划
func (s fbaShipmentService) Plans(params FBAShipmentPlansQueryParams) (items []FBAShipmentPlan, nextOffset int, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.SetPagingVars()
	res := struct {
		NormalResponse
		Data struct {
			PlanList []FBAShipmentPlan `json:"plan_list"` // 发货计划组列表
		} `json:"data"`
	}{}
	resp, err := s.httpClient.R().
		SetBody(params).
		Post("/data/fba_report/shipmentPlanLists")
	if err != nil {
		return
	}

	if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
		items = res.Data.PlanList
		nextOffset = params.NextOffset
		isLastPage = len(items) < params.Limit
	}
	return
}
