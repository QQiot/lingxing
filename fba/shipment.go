package fba

// FBA发货单列表

// ShipmentLogistics 物流
type ShipmentLogistics struct {
	ReplaceTrackingNumber string `json:"replace_tracking_number"` // 跟踪单号
	TrackingNumber        string `json:"tracking_number"`         // 物流商号
}

// Shipment 货件
type Shipment struct {
	ID                             int      `json:"id"`                                // 明细 ID
	MID                            int      `json:"mid"`                               // 国家 ID
	DestinationFulfillmentCenterId int      `json:"destination_fulfillment_center_id"` // 物流中心编码
	QuantityShipped                int      `json:"quantity_shipped"`                  // 申报量
	WarehouseName                  string   `json:"wname"`                             // 仓库名称
	ShipmentSN                     string   `json:"shipment_sn"`                       // 发货单号
	ShipmentId                     string   `json:"shipment_id"`                       // 货件id
	Wid                            string   `json:"wid"`                               // 仓库id
	Pid                            int      `json:"pid"`                               // 货件明细 ID
	SellerName                     string   `json:"sname"`                             // 店铺名称
	ProductName                    string   `json:"product_name"`                      // 产品名称
	Num                            string   `json:"num"`                               // 发货数量
	PicURL                         string   `json:"pic_url"`                           // 图片 URL
	PackingType                    int      `json:"packing_type"`                      // 混装类型 2原装 1原装
	FulfillmentNetworkSKU          string   `json:"fulfillment_network_sku"`           // listing的fnsku
	SKU                            string   `json:"sku"`                               // sku
	FnSKU                          string   `json:"fnsku"`                             // 仓库fnsku
	MSKU                           string   `json:"msku"`                              // seller_sku
	Nation                         string   `json:"nation"`                            // 国家名称
	ApplyNum                       string   `json:"apply_num"`                         // 关联货件量
	ProductId                      string   `json:"product_id"`                        // 商品 ID
	Remark                         string   `json:"remark"`                            // 备注
	Status                         int      `json:"status"`                            // 状态
	SID                            int      `json:"sid"`                               // 店铺 ID
	IsCombo                        bool     `json:"is_combo"`                          // 是否组合商品
	CreateByMWS                    int      `json:"create_by_mws"`                     // 创建发货单的途径
	WhbCodeList                    int      `json:"whb_code_list"`                     // 仓位编码列表
	PackingTypeName                int      `json:"packing_type_name"`                 // 包装名称
	ProductValidNum                int      `json:"product_valid_num"`                 // 可用量
	ProductQcNum                   int      `json:"product_qc_num"`                    // 待检量
	DiffNum                        int      `json:"diff_num"`                          // 差额
	NotRelateList                  []string `json:"not_relate_list"`                   // 未关联货件列表
	StatusName                     string   `json:"status_name"`                       // 状态名称
	HeadFeeTypeName                string   `json:"head_fee_type_name"`                // 头程分摊名称
	FileList                       string   `json:"fileList"`                          // 文件列表
}

type ShipmentSheet struct {
	ID                   int                 `json:"id"`                     // 发货单 ID
	ShipmentSN           string              `json:"shipment_sn"`            // 发货单号
	Status               int                 `json:"status"`                 // 发货单状态（-1：待配货、0：待发货、1：已发货、2：已完成、3：已作废）
	ShipmentTime         string              `json:"shipment_time"`          // 发货时间
	WarehouseName        string              `json:"wname"`                  // 仓库名称
	CreateUser           string              `json:"create_user"`            // 创建用户
	LogisticsChannelName string              `json:"logistics_channel_name"` // 物流方式
	ExpectedArrivalDate  string              `json:"expected_arrival_date"`  // 到货时间
	ETDDate              string              `json:"etd_date"`               // 开船时间
	ETADate              string              `json:"eta_date"`               // 预计到港时间
	DeliveryDate         string              `json:"delivery_date"`          // 实际妥投时间
	CreateTime           string              `json:"create_time"`            // 创建时间
	IsPick               bool                `json:"is_pick"`                // 拣货状态（0：未拣货、1：已拣货）
	IsPrint              bool                `json:"is_print"`               // 是否打印
	PickTime             string              `json:"pick_time"`              // 拣货时间
	PrintNum             string              `json:"print_num"`              // 打印次数
	HeadFeeType          int                 `json:"head_fee_type"`          // 头程费分配方式（0：按计费重、1：按实重、2：按体积重、3：按SKU数量、4：自定义、5：按箱子体积）
	FileId               string              `json:"file_id"`                // 附件文件
	GMTModified          string              `json:"gmt_modified"`           // 更新时间
	Remark               string              `json:"remark"`                 // 备注
	WarehouseId          int                 `json:"wid"`                    // 仓库 ID
	IsReturnStock        bool                `json:"is_return_stock"`        // 是否恢复库存
	Logistics            []ShipmentLogistics `json:"logistics"`              // 物流列表
	Relatelist           []Shipment          `json:"relate_list"`            // 关联货件列表
}

type ShipmentSheetsQueryParams struct {
	SearchValue   string   `json:"search_value,omitempty"`   // 搜索的值
	SearchField   string   `json:"search_field,omitempty"`   // 搜索字段（shipment_sn：发货单号、sku：SKU、shipment_id：货件单号）
	SIDs          []string `json:"sids,omitempty"`           // 店铺id
	MIDs          []string `json:"mids,omitempty"`           // 国家id
	WIDs          []string `json:"wid,omitempty"`            // 仓库id
	LogisticsType []string `json:"logistics_type,omitempty"` // 物流方式id
	Status        int      `json:"status,omitempty"`         // 发货单状态（-1：待配货、0：待发货、1：已发货、3：已作废）
	PrintStatus   string   `json:"print_status,omitempty"`   // 打印状态（0：未打印、1：已打印）
	PickStatus    string   `json:"pick_status,omitempty"`    // 拣货状态（0：未拣货、1：已拣货）
	TimeType      int      `json:"time_type,omitempty"`      // 按时间查询时必传时间类型（ 0：发货时间、1：到货时间、2：创建时间 ）
	StartDate     string   `json:"start_date,omitempty"`     // 开始日期
	EndDate       string   `json:"end_date,omitempty"`       // 结束日期
}
