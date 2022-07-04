package entity

type Email struct {
	WebMailUUID   string `json:"webmail_uuid"`   // 邮件唯一标识
	Date          string `json:"date"`           // 日期
	Subject       string `json:"subject"`        // 邮件标题
	FromName      string `json:"from_name"`      // 发件人姓名
	FromAddress   string `json:"from_address"`   // 发件人地址
	ToName        string `json:"to_name"`        // 接收人
	ToAddress     string `json:"to_address"`     // 接收人地址
	HasAttachment int    `json:"has_attachment"` // 是否存在附件（0：不存在、1：存在）
}
