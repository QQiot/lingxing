package config

type Config struct {
	Debug     bool   // 是否启用调试模式
	Sandbox   bool   // 是否为沙箱环境
	AppId     string // APP ID
	AppSecret string // APP Secret
}
