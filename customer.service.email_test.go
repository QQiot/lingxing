package lingxing

import (
	"github.com/hiscaler/gox/jsonx"
	"testing"
)

func TestCustomerServiceEmailService_All(t *testing.T) {
	params := CustomerServiceEmailsQueryParams{
		Flag:  "receive",
		Email: "1@gmail.com",
	}
	params.Limit = 1
	var emails []Email
	for {
		items, nextOffset, isLastPage, err := lingXingClient.Services.CustomerService.Email.All(params)
		if err != nil {
			t.Errorf("Services.CustomerService.Email.All() error: %s", err.Error())
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
