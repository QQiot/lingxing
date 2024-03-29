package lingxing

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	jsoniter "github.com/json-iterator/go"
)

type customerServiceEmailService service

// 邮件列表

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

type CustomerServiceEmailsQueryParams struct {
	Paging
	Flag  string `json:"flag" url:"flag"`   // 类型
	Email string `json:"email" url:"email"` // 邮箱
}

func (m CustomerServiceEmailsQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Flag,
			validation.Required.Error("类型不能为空"),
			validation.In("sent", "receive").Error("无效的类型"),
		),
		validation.Field(&m.Email,
			validation.Required.Error("邮箱不能为空"),
			is.EmailFormat.Error("无效的邮箱格式"),
		),
	)
}

// All 邮件列表
// https://openapidoc.lingxing.com/#/docs/Service/lists
func (s customerServiceEmailService) All(params CustomerServiceEmailsQueryParams) (items []Email, nextOffset int, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.SetPagingVars()
	res := struct {
		NormalResponse
		Data []Email `json:"data"`
	}{}
	resp, err := s.httpClient.R().
		SetQueryParamsFromValues(toValues(params)).
		Get("/data/mail/lists")
	if err != nil {
		return
	}

	if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
		items = res.Data
		nextOffset = params.nextOffset
		isLastPage = len(items) < params.Limit
	}
	return
}

//  邮件详情
// https://openapidoc.lingxing.com/#/docs/Service/detail

type CustomerServiceEmailAttachment struct {
	Name string `json:"name"` // 附件名称
	Size int    `json:"size"` // 附件大小（byte）
}

type CustomerServiceEmail struct {
	WebMailUUID  string                           `json:"webmail_uuid"`   // 邮件唯一标识
	Subject      string                           `json:"subject"`        // 邮件标题
	FromName     string                           `json:"from_name"`      // 发件人姓名
	FromAddress  string                           `json:"from_address"`   // 发件人地址
	ToAddressAll string                           `json:"to_address_all"` // 所有收件人地址
	Date         string                           `json:"date"`           // 日期
	CC           string                           `json:"cc"`             // 抄送
	BCC          string                           `json:"bcc"`            // 密送地址
	TextHtml     string                           `json:"text_html"`      // 邮件内容
	Attachments  []CustomerServiceEmailAttachment `json:"attachments"`    // 附件
}

func (s customerServiceEmailService) One(webMailUUID string) (item CustomerServiceEmail, err error) {
	res := struct {
		NormalResponse
		Data CustomerServiceEmail `json:"data"`
	}{}
	resp, err := s.httpClient.R().
		SetBody(map[string]string{"webmail_uuid": webMailUUID}).
		Get("/data/mail/detail")
	if err != nil {
		return
	}

	if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
		item = res.Data
	}
	return
}
