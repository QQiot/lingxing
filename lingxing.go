package lingxing

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/hiscaler/gox/cryptox"
	"github.com/hiscaler/gox/jsonx"
	"github.com/hiscaler/gox/stringx"
	"github.com/hiscaler/lingxing/config"
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
	"github.com/spf13/cast"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// https://openapidoc.lingxing.com/#/docs/Guidance/ErrorCode
const (
	OK                       = 200     // 无错误
	AppIdNotExistError       = 2001001 // appId 不存在
	InvalidAppSecretError    = 2001002 // appSecret 不正确或者 urlencode 需要进行编码
	AccessTokenExpireError   = 2001003 // token 不存在或者已经过期
	UnauthorizedError        = 2001004 // api 未授权
	InvalidAccessTokenError  = 2001005 // token 不正确
	SignError                = 2001006 // 签名错误
	SignExpiredError         = 2001007 // 签名过期
	RefreshTokenExpiredError = 2001008 // RefreshToken 过期
	InvalidRefreshTokenError = 2001009 // 无效的 RefreshToken
	InvalidQueryParamsError  = 3001001 // 查询参数缺失
	InvalidIPError           = 3001002 // 应用所在服务器的 ip 不在白名单中
	TooManyRequestsError     = 3001008 // 接口请求超请求次数限额
)

var ErrNotFound = errors.New("lingxing: not found")

type defaultQueryParams struct {
	Offset   int // 分页偏移索引（默认0）
	Limit    int // 分页偏移长度（默认1000）
	MaxLimit int // 最大偏移长度
}

type LingXing struct {
	host               string
	appId              string
	appSecret          string
	accessToken        string
	Debug              bool               // 是否调试模式
	Client             *resty.Client      // HTTP 客户端
	MerchantId         string             // 商户 ID
	Logger             *log.Logger        // 日志
	DefaultQueryParams defaultQueryParams // 查询默认值
	auth               AuthResponse
}

func init() {
	extra.RegisterFuzzyDecoders()
}

func NewLingXing(config config.Config) *LingXing {
	logger := log.New(os.Stdout, "[ LingXing ] ", log.LstdFlags|log.Llongfile)
	lx := &LingXing{
		appId:     config.AppId,
		appSecret: config.AppSecret,
		Logger:    logger,
		DefaultQueryParams: defaultQueryParams{
			Offset:   0,
			Limit:    1000,
			MaxLimit: 1000,
		},
	}
	client := resty.New().SetDebug(config.Debug).
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
		})
	if config.Debug {
		client.SetBaseURL("https://openapisandbox.lingxing.com/erp/sc")
	} else {
		client.SetBaseURL("https://openapi.lingxing.com/erp/sc")
	}

	client.SetTimeout(10 * time.Second).
		OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
			if lx.auth.ExpiresIn <= 0 || lx.auth.AccessToken == "" {
				if auth, err := lx.Auth(lx.appId, lx.appSecret, config.Debug); err != nil {
					logger.Printf("auth error: %s", err.Error())
					return err
				} else {
					lx.auth = auth
				}
			}

			queryParams := map[string]string{
				"app_key":      config.AppId,
				"access_token": lx.auth.AccessToken,
				"timestamp":    strconv.FormatInt(time.Now().Unix(), 10),
			}
			params := cast.ToStringMap(jsonx.ToJson(request.Body, "{}")) // Body
			if params == nil {
				params = make(map[string]interface{}, 3)
			}
			for k, v := range queryParams {
				params[k] = v
			}
			sign, err := lx.generateSign(params)
			if err != nil {
				return err
			}

			queryParams["sign"] = url.QueryEscape(sign)
			request.SetQueryParams(queryParams)
			return nil
		}).
		OnAfterResponse(func(client *resty.Client, response *resty.Response) (err error) {
			if response.IsSuccess() {
				r := struct {
					Code    interface{} `json:"code"`
					Message string      `json:"msg"`
				}{}
				if err = jsoniter.Unmarshal(response.Body(), &r); err == nil {
					err = ErrorWrap(cast.ToInt(r.Code), r.Message)
				}
			}
			if err != nil {
				logger.Printf("OnAfterResponse error: %s", err.Error())
			}
			return
		}).
		SetRetryCount(2).
		SetRetryWaitTime(5 * time.Second).
		SetRetryMaxWaitTime(10 * time.Second).
		AddRetryCondition(func(response *resty.Response, err error) bool {
			if response == nil {
				return false
			}

			retry := response.StatusCode() == http.StatusTooManyRequests
			if !retry {
				r := struct{ Code int }{}
				retry = jsoniter.Unmarshal(response.Body(), &r) == nil && r.Code == TooManyRequestsError
			}
			if retry {
				text := response.Request.URL
				if err != nil {
					text += fmt.Sprintf(", error: %s", err.Error())
				}
				logger.Printf("Retry request: %s", text)
			}
			return retry
		})
	lx.Client = client
	return lx
}

type NormalResponse struct {
	Code         int         `json:"code"`
	Message      string      `json:"message"`
	ErrorDetails interface{} `json:"error_details"`
	RequestId    string      `json:"request_id"`
	ResponseTime string      `json:"response_time"`
	Data         interface{} `json:"data"`
	Total        int         `json:"total"`
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

func (lx *LingXing) Auth(appId, appSecret string, debug bool) (ar AuthResponse, err error) {
	result := struct {
		Code    string       `json:"code"`
		Message string       `json:"msg"`
		Data    AuthResponse `json:"data"`
	}{}

	client := resty.New().
		SetDebug(true).
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
		})
	if debug {
		client.SetBaseURL("https://openapisandbox.lingxing.com")
	} else {
		client.SetBaseURL("https://openapi.lingxing.com")
	}
	resp, err := client.R().
		SetResult(&result).
		Post(fmt.Sprintf("/api/auth-server/oauth/access-token?appId=%s&appSecret=%s", appId, url.QueryEscape(appSecret)))
	if err != nil {
		return
	}

	code, _ := strconv.ParseInt(result.Code, 10, 32)
	if resp.IsSuccess() {
		err = ErrorWrap(int(code), result.Message)
		if err == nil {
			ar = result.Data
		}
	} else {
		if e := jsoniter.Unmarshal(resp.Body(), &result); e == nil {
			err = ErrorWrap(int(code), result.Message)
		} else {
			err = errors.New(resp.Status())
		}
	}
	return
}

func (lx *LingXing) generateSign(params map[string]interface{}) (sign string, err error) {
	n := len(params)
	keys := make([]string, n)
	i := 0
	for k := range params {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	sb := strings.Builder{}
	for _, key := range keys {
		sb.WriteString(key)
		sb.WriteRune('=')
		switch v := params[key].(type) {
		case string:
			sb.WriteString(v)
		default:
			var b []byte
			b, err = jsoniter.Marshal(v)
			if err == nil {
				sb.Write(b)
			} else {
				return
			}
		}
		sb.WriteRune('&')
	}
	s := sb.String()
	if n = len(s); n > 0 {
		s = s[0 : n-1]
	}

	aesTool := NewAesTool(stringx.ToBytes(lx.appId), len(lx.appId))
	aesEncrypted, err := aesTool.ECBEncrypt(stringx.ToBytes(strings.ToUpper(cryptox.Md5(s))))
	if err != nil {
		return
	}

	sign = base64.StdEncoding.EncodeToString(aesEncrypted)
	return
}

// ErrorWrap 错误包装
func ErrorWrap(code int, message string) error {
	if code == OK || code == 0 {
		return nil
	}

	message = strings.TrimSpace(message)
	if message == "" {
		switch code {
		case AppIdNotExistError:
			message = "appId 不存在"
		case InvalidAppSecretError:
			message = "appSecret 不正确或者未编码"
		case AccessTokenExpireError:
			message = "token 不存在或者已经过期"
		case UnauthorizedError:
			message = "api 未授权"
		case InvalidAccessTokenError:
			message = "token 不正确"
		case SignError:
			message = "签名错误"
		case SignExpiredError:
			message = "签名过期"
		case RefreshTokenExpiredError:
			message = "RefreshToken 过期"
		case InvalidRefreshTokenError:
			message = "无效的 RefreshToken"
		case InvalidQueryParamsError:
			message = "查询参数缺失"
		case InvalidIPError:
			message = "应用所在服务器的 ip 不在白名单中"
		case TooManyRequestsError:
			message = "接口请求超请求次数限额"
		default:
			message = "未知错误"
		}
	}
	return fmt.Errorf("%d: %s", code, message)
}
