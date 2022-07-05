package lingxing

import (
	"github.com/hiscaler/gox/jsonx"
	"github.com/hiscaler/lingxing/entity"
	"testing"
)

func TestCustomerServiceEmailService_All(t *testing.T) {
	params := CustomerServiceEmailsQueryParams{
		Flag:  "1",
		Email: "1@gmail.com",
	}
	params.Limit = 1
	var emails []entity.Email
	for {
		items, nextOffset, isLastPage, err := lingXingClient.Services.CustomerService.Email.All(params)
		if err != nil {
			t.Errorf("Services.CustomerService.Emails() error: %s", err.Error())
		} else {
			emails = append(emails, items...)
		}
		if isLastPage || err != nil {
			break
		}
		params.Offset = nextOffset
	}
	t.Log(jsonx.ToPrettyJson(emails))
}
