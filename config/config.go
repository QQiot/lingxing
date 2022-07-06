package config

type Config struct {
	Debug       bool // 是否启用调试模式
	Sandbox     bool // 是否为沙箱环境（true 启用 Dev 配置，false 启用 Prod 配置）
	Environment struct {
		Dev struct {
			AppId     string // APP ID
			AppSecret string // APP Secret
		} // 开发环境配置
		Prod struct {
			AppId     string // APP ID
			AppSecret string // APP Secret
		} // 生产环境配置
	} // 环境配置
	EnableCache bool // 是否激活缓存
}
