package lingxing

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/hiscaler/lingxing/config"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	OK                       = 200     // 无错误
	AppIdNotExistError       = 2001001 // appId 不存在
	InvalidAppSecretError    = 2001002 // appSecret 不正确或者 urlencode 需要进行编码
	AccessTokenExpireError   = 2001003 // token不存在或者已经过期
	UnauthorizedError        = 2001004 // api未授权
	InvalidAccessTokenError  = 2001005 // token不正确
	SignError                = 2001006 // 签名错误
	SignExpiredError         = 2001007 // 签名过期
	RefreshTokenExpiredError = 2001008 // RefreshToken 过期
	InvalidRefreshTokenError = 2001009 // 无效的 RefreshToken
	InvalidQueryParamsError  = 3001001 // 查询参数缺失
	InvalidIPError           = 3001002 // IP 不在白名单内
	TooManyRequestsError     = 3001008 // 接口请求超请求次数限额
)

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
	client := resty.New().SetDebug(true).
		SetBaseURL("https://openapi.lingxing.com/erp/sc").
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
		}).
		SetTimeout(10 * time.Second).
		OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
			timestamp := time.Now().Unix()
			if lx.auth.ExpiresIn > timestamp {
				if auth, err := lx.Auth(lx.appId, lx.appSecret); err != nil {
					logger.Printf("auth error: %s", err.Error())
					return err
				} else {
					lx.auth = auth
				}
			}
			sign, err := lx.generateSign(map[string]interface{}{
				"app_key":      config.AppId,
				"access_token": lx.accessToken,
				"timestamp":    timestamp,
			})
			if err != nil {
				return err
			}
			qp := map[string]string{
				"app_key":      config.AppId,
				"sign":         sign,
				"access_token": lx.accessToken,
				"timestamp":    strconv.FormatInt(timestamp, 10),
			}
			request.SetQueryParams(qp)
			return nil
		}).
		OnAfterResponse(func(client *resty.Client, response *resty.Response) (err error) {
			if response.IsSuccess() {
				r := struct {
					Code    string `json:"code"`
					Message string `json:"message"`
				}{}
				if err = json.Unmarshal(response.Body(), &r); err != nil {
					return
				}
				err = ErrorWrap(r.Code, r.Message)
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
				retry = json.Unmarshal(response.Body(), &r) == nil && r.Code == TooManyRequestsError
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
	Code         string      `json:"code"`
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
	ExpiresIn    int64  `json:"expires_in"`
}

func (lx *LingXing) Auth(appId, appSecret string) (ar AuthResponse, err error) {
	result := struct {
		NormalResponse
		Data AuthResponse `json:"data"`
	}{}
	resp, err := lx.Client.R().
		SetResult(&result).
		Post(fmt.Sprintf("/api/auth-server/oauth/access-token?appId=%s&appSecret=%s", appId, appSecret))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = ErrorWrap(result.Code, result.Message)
		if err == nil {
			ar = result.Data
		}
	} else {
		if e := json.Unmarshal(resp.Body(), &result); e == nil {
			err = ErrorWrap(result.Code, result.Message)
		} else {
			err = errors.New(resp.Status())
		}
	}
	return
}

func (lx *LingXing) generateSign(params map[string]interface{}) (sign string, err error) {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var qStrList []string
	for _, key := range keys {
		switch v := params[key].(type) {
		case string:
			qStrList = append(qStrList, fmt.Sprintf("%s=%s", key, v))
		default:
			var jsonV []byte
			jsonV, err = json.Marshal(v)
			if err != nil {
				return
			}
			qStrList = append(qStrList, fmt.Sprintf("%s=%s", key, string(jsonV)))
		}
	}

	md5Str := strings.ToUpper(fmt.Sprintf("%x", md5.Sum([]byte(strings.Join(qStrList, "&")))))
	key := lx.appId
	aesTool := NewAesTool([]byte(key), len(key))
	aesEncrypted, err := aesTool.ECBEncrypt([]byte(md5Str))
	if err != nil {
		return
	}

	sign = base64.StdEncoding.EncodeToString(aesEncrypted)
	return
}

// ErrorWrap 错误包装
func ErrorWrap(code string, message string) error {
	c, _ := strconv.Atoi(code)
	if c == OK {
		return nil
	}

	message = strings.TrimSpace(message)
	if message == "" {
		switch c {
		case AppIdNotExistError:
			message = "appId 不存在"
		case InvalidAppSecretError:
			message = "appSecret 不正确或者 urlencode 需要进行编码"
		case AccessTokenExpireError:
			message = "token不存在或者已经过期"
		case UnauthorizedError:
			message = "api未授权"
		case InvalidAccessTokenError:
			message = "token不正确"
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
			message = "IP 不在白名单内"
		case TooManyRequestsError:
			message = "接口请求超请求次数限额"
		default:
			message = "未知错误"
		}
	}
	return fmt.Errorf("%s: %s", code, message)
}
