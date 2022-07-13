package lingxing

import (
	"testing"
)

func TestAuthorizationService_Get(t *testing.T) {
	ar, err := lingXingClient.SetDebug(false).Services.Authorization.GetToken()
	if err != nil {
		t.Errorf("get token error: %s", err.Error())
	} else {
		t.Logf("accessToken: %+v", ar)
		ar, err = lingXingClient.Services.Authorization.RefreshToken(ar.RefreshToken)
		if err != nil {
			t.Errorf("get token error: %s", err.Error())
		} else {
			t.Logf("RefreshToken: %+v", ar)
		}
	}
}
