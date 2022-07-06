package entity

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
	IsMaster      bool   `json:"is_master"`       // 是否为主账号
}
