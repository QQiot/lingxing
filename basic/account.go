package basic

import (
	"encoding/json"
	"errors"
	"github.com/hiscaler/lingxing"
	"time"
)

// 支持查询企业开启的全部ERP账号
// https://openapidoc.lingxing.com/#/docs/BasicData/AccoutLists

type Account struct {
	UID           string    `json:"uid"`             // 账号ID
	RealName      string    `json:"realname"`        // 姓名
	Username      string    `json:"username"`        // 用户名
	Mobile        string    `json:"mobile"`          // 电话
	Email         string    `json:"email"`           // 邮箱
	LoginNum      string    `json:"login_num"`       // 登陆次数
	LastLoginTime time.Time `json:"last_login_time"` // 最近登录时间
	LastLoginIP   string    `json:"last_login_ip"`   // 最近登录IP
	Status        string    `json:"status"`          // 状态（0：禁用、1：正常）
	CreateTime    time.Time `json:"create_time"`     // 创建时间
	Role          string    `json:"role"`            // 角色
	Seller        string    `json:"seller"`          // 店铺权限
	IsMaster      bool      `json:"is_master"`       // 是否为主账号
}

func (s service) Accounts() (items []Account, err error) {
	res := struct {
		lingxing.NormalResponse
		Data []Account `json:"data"`
	}{}
	resp, err := s.lingXing.Client.R().
		SetResult(&res).
		Post("/data/account/lists")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = lingxing.ErrorWrap(res.Code, res.Message)
		if err == nil {
			items = res.Data
		}
	} else {
		if e := json.Unmarshal(resp.Body(), &res); e == nil {
			err = lingxing.ErrorWrap(res.Code, res.Message)
		} else {
			err = errors.New(resp.Status())
		}
	}
	return
}
