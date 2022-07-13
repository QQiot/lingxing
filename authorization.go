package lingxing

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/hiscaler/gox/bytex"
	"net/url"
	"strconv"
	"time"
)

type authorizationService service

func newHttpClient(c cfg) *resty.Client {
	httpClient := resty.New().SetDebug(c.debug).
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
			"User-Agent":   userAgent,
		})
	if c.sandbox {
		httpClient.SetBaseURL("https://openapisandbox.lingxing.com")
	} else {
		httpClient.SetBaseURL("https://openapi.lingxing.com")
	}
	return httpClient
}

// GetToken 获取 access-token 和 refresh-token
// https://openapidoc.lingxing.com/#/docs/Authorization/GetToken
func (s authorizationService) GetToken() (ar authorizationResponse, err error) {
	result := struct {
		Code    string                `json:"code"`
		Message string                `json:"msg"`
		Data    authorizationResponse `json:"data"`
	}{}
	resp, err := newHttpClient(*s.config).R().
		SetResult(&result).
		Post(fmt.Sprintf("/api/auth-server/oauth/access-token?appId=%s&appSecret=%s", s.config.appId, url.QueryEscape(s.config.appSecret)))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		code, _ := strconv.Atoi(result.Code)
		if err = ErrorWrap(code, result.Message); err == nil {
			ar = result.Data
			ar.ExpiresDatetime = time.Now().Add(time.Duration(ar.ExpiresIn*5/10) * time.Second) // 剩余 1/2 时间就会要求更换 token
		}
	} else {
		err = fmt.Errorf("%s: %s", resp.Status(), bytex.ToString(resp.Body()))
	}
	return
}

// RefreshToken 刷新token（token续约，每个refreshToken只能用一次）
// https://openapidoc.lingxing.com/#/docs/Authorization/RefreshToken
func (s authorizationService) RefreshToken(refreshToken string) (ar authorizationResponse, err error) {
	result := struct {
		Code    string                `json:"code"`
		Message string                `json:"msg"`
		Data    authorizationResponse `json:"data"`
	}{}
	resp, err := newHttpClient(*s.config).R().
		SetResult(&result).
		Post(fmt.Sprintf("/api/auth-server/oauth/refresh?appId=%s&refreshToken=%s", s.config.appId, refreshToken))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		code, _ := strconv.Atoi(result.Code)
		if err = ErrorWrap(code, result.Message); err == nil {
			ar = result.Data
		}
	} else {
		err = fmt.Errorf("%s: %s", resp.Status(), bytex.ToString(resp.Body()))
	}
	return
}
