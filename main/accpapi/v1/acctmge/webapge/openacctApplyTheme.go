package webapge

type OpenacctApplyTheme struct {
	// 基础色。默认值为#1976D2 wap端有效产品类型。
	Primary string `json:"primary"`
	// 间色。默认值为#424242 wap端有效。
	Secondary string `json:"secondary"`
	// 强调色。默认值为#82B1FF wap端有效。
	Accent string `json:"accent"`
	// 错误色。默认值为#FF5252 wap端有效。
	Error string `json:"error"`
	// 信息色。默认值为#2196F3 wap端有效。
	Info string `json:"info"`
	// 成功色。默认值为#4CAF50 wap端有效。
	Success string `json:"success"`
	// 警告色。默认值为#FFC107 wap端有效。
	Warning string `json:"warning"`
}
