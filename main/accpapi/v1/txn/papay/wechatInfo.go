package papay

type WechatInfo struct {
	// 商品描述，商品或支付单简要描述。注：微信付款码支付时为必输。
	Body string `json:"body"`
	// 商品详情。
	Detail string `json:"detail"`
}
