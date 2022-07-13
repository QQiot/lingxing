package lingxing

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
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
		err = errors.New(resp.Status())
	}
	return
}

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
		err = errors.New(resp.Status())
	}
	return
}
