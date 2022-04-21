package config

type Config struct {
	Debug       bool   // 是否为调试模式
	AppId       string // APP ID
	AppSecret   string // APP Secret
	EnableCache bool   // 是否激活缓存
}
