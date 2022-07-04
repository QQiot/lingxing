package entity

import "time"

type Rate struct {
	Date       string    `json:"date"`        // 汇率年月
	Code       string    `json:"code"`        // 币种
	Icon       string    `json:"icon"`        // 币种符号
	Name       string    `json:"name"`        // 币种名
	RateOrg    float64   `json:"rate_org"`    // 官方汇率（数据来源中国银行官方汇率）
	MyRate     float64   `json:"my_rate"`     // 我的汇率（用户自定义汇率，系统首先使用该汇率数据）
	UpdateTime time.Time `json:"update_time"` // 更新时间
}
