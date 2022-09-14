package config

type Config struct {
	Debug     bool   // 是否启用调试模式
	Timeout   int    // HTTP 超时设定（单位：秒）
	Sandbox   bool   // 是否为沙箱环境
	AppId     string // APP ID
	AppSecret string // APP Secret
}
