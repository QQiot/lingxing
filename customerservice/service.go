package customerservice

import (
	"github.com/hiscaler/lingxing"
)

type service struct {
	lingXing *lingxing.LingXing
}

type Service interface {
	Emails(params EmailsQueryParams) (items []Email, nextOffset int, isLastPage bool, err error) // 邮件列表
	Email(webMailUUID string) (item EmailDetail, err error)                                      // 邮件详情
}

func NewService(lx *lingxing.LingXing) Service {
	return service{lx}
}
