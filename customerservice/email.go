package customerservice

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hiscaler/lingxing"
	jsoniter "github.com/json-iterator/go"
)

// 邮件列表
// https://openapidoc.lingxing.com/#/docs/Service/lists

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

type EmailsQueryParams struct {
	lingxing.Paging
	Flag  string `json:"flag"`  // 类型
	Email string `json:"email"` // 邮箱
}

func (m EmailsQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Flag, validation.Required.Error("类型不能为空")),
		validation.Field(&m.Email, validation.Required.Error("邮箱不能为空")),
	)
}

func (s service) Emails(params EmailsQueryParams) (items []Email, nextOffset int, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.SetPagingVars(params.Offset, params.Limit, s.lingXing.DefaultQueryParams.MaxLimit)
	res := struct {
		lingxing.NormalResponse
		Data []Email `json:"data"`
	}{}
	resp, err := s.lingXing.Client.R().
		SetBody(params).
		Get("/data/mail/lists")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
			if err = lingxing.ErrorWrap(res.Code, res.Message); err == nil {
				items = res.Data
				nextOffset = params.NextOffset
				isLastPage = len(items) < params.Limit
			}
		}
	} else {
		if e := jsoniter.Unmarshal(resp.Body(), &res); e == nil {
			err = lingxing.ErrorWrap(res.Code, res.Message)
		} else {
			err = errors.New(resp.Status())
		}
	}
	if err != nil {
		return
	}

	return
}

//  邮件详情
// https://openapidoc.lingxing.com/#/docs/Service/detail

type EmailAttachment struct {
	Name string `json:"name"` // 附件名称
	Size int    `json:"size"` // 附件大小（b）
}

type EmailDetail struct {
	WebMailUUID  string            `json:"webmail_uuid"`   // 邮件唯一标识
	Subject      string            `json:"subject"`        // 邮件标题
	FromName     string            `json:"from_name"`      // 发件人姓名
	FromAddress  string            `json:"from_address"`   // 发件人地址
	ToAddressAll string            `json:"to_address_all"` // 所有收件人地址
	Date         string            `json:"date"`           // 日期
	CC           string            `json:"cc"`             // 抄送
	BCC          string            `json:"bcc"`            // 密送地址
	TextHtml     string            `json:"text_html"`      // 邮件内容
	Attachments  []EmailAttachment `json:"attachments"`    // 附件
}

func (s service) Email(webMailUUID string) (item EmailDetail, err error) {
	res := struct {
		lingxing.NormalResponse
		Data EmailDetail `json:"data"`
	}{}
	resp, err := s.lingXing.Client.R().
		SetBody(map[string]string{"webmail_uuid": webMailUUID}).
		Get("/data/mail/detail")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
			if err = lingxing.ErrorWrap(res.Code, res.Message); err == nil {
				item = res.Data
			}
		}
	} else {
		if e := jsoniter.Unmarshal(resp.Body(), &res); e == nil {
			err = lingxing.ErrorWrap(res.Code, res.Message)
		} else {
			err = errors.New(resp.Status())
		}
	}
	return
}
