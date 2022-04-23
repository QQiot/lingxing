package basic

import (
	"errors"
	"github.com/hiscaler/lingxing"
	jsoniter "github.com/json-iterator/go"
)

// 支持查询企业开启的全部ERP账号
// https://openapidoc.lingxing.com/#/docs/BasicData/AccoutLists

type Account struct {
	UID           int    `json:"uid"`             // 账号ID
	RealName      string `json:"realname"`        // 姓名
	Username      string `json:"username"`        // 用户名
	Mobile        string `json:"mobile"`          // 电话
	Email         string `json:"email"`           // 邮箱
	LoginNum      int    `json:"login_num"`       // 登录次数
	LastLoginTime string `json:"last_login_time"` // 最近登录时间
	LastLoginIP   string `json:"last_login_ip"`   // 最近登录IP
	Status        int    `json:"status"`          // 状态（0：禁用、1：正常）
	CreateTime    string `json:"create_time"`     // 创建时间
	ZID           int    `json:"zid"`             // ?
	Role          string `json:"role"`            // 角色
	Seller        string `json:"seller"`          // 店铺权限
	IsMaster      int    `json:"is_master"`       // 是否为主账号
}

func (s service) Accounts() (items []Account, err error) {
	res := struct {
		lingxing.NormalResponse
		Data []Account `json:"data"`
	}{}
	resp, err := s.lingXing.Client.R().Get("/data/account/lists")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
			if err = lingxing.ErrorWrap(res.Code, res.Message); err == nil {
				items = res.Data
			}
		}
	} else {
		if e := jsoniter.Unmarshal(resp.Body(), &res); e == nil {
			err = lingxing.ErrorWrap(res.Code, res.Message)
		} else {
			err = errors.New(resp.Status())
		}
	}
	return
}
