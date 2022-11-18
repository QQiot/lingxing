package lingxing

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/hiscaler/gox/bytex"
	"github.com/hiscaler/lingxing/config"
	"net/url"
	"strconv"
	"time"
)

type authorizationService service

func newHttpClient(c config.Config) *resty.Client {
	httpClient := resty.
		New().
		SetDebug(c.Debug).
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
			"User-Agent":   userAgent,
		})
	if c.Sandbox {
		httpClient.SetBaseURL("https://openapisandbox.lingxing.com")
	} else {
		httpClient.SetBaseURL("https://openapi.lingxing.com")
	}
	return httpClient
}

// GetToken 获取 access-token 和 refresh-token
// https://openapidoc.lingxing.com/#/docs/Authorization/GetToken
func (s authorizationService) GetToken() (ar Token, err error) {
	result := struct {
		Code    string `json:"code"`
		Message string `json:"msg"`
		Data    Token  `json:"data"`
	}{}
	resp, err := newHttpClient(*s.config).
		SetLogger(s.logger).
		R().
		SetResult(&result).
		Post(fmt.Sprintf("/api/auth-server/oauth/access-token?appId=%s&appSecret=%s", s.config.AppId, url.QueryEscape(s.config.AppSecret)))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		code, _ := strconv.Atoi(result.Code)
		if err = ErrorWrap(code, result.Message); err == nil {
			ar = result.Data
			ar.ExpiresDatetime = time.Now().Unix() + int64(ar.ExpiresIn*4/5) // 剩余 1/5 时间需要更换 token
		}
	} else {
		err = fmt.Errorf("%s: %s", resp.Status(), bytex.ToString(resp.Body()))
	}
	return
}

// RefreshToken 刷新 token（token 续约，每个 refreshToken 只能用一次）
// https://openapidoc.lingxing.com/#/docs/Authorization/RefreshToken
func (s authorizationService) RefreshToken(refreshToken string) (ar Token, err error) {
	result := struct {
		Code    string `json:"code"`
		Message string `json:"msg"`
		Data    Token  `json:"data"`
	}{}
	resp, err := newHttpClient(*s.config).
		SetLogger(s.logger).
		R().
		SetResult(&result).
		Post(fmt.Sprintf("/api/auth-server/oauth/refresh?appId=%s&refreshToken=%s", s.config.AppId, refreshToken))
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
