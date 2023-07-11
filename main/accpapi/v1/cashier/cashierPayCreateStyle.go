package cashier

type CashierPayCreateStyle struct {
	// 支付页面的背景颜色。 可选范围为 000000 ~ ffffff， 默认值为ff5001
	BgColor string `json:"bg_color"`
	// 支付页面字体颜色。可选范围为 000000 ~ ffffff。
	FontColor string `json:"font_color"`
}
