package txn

/**
 * 网关类支付 响应参数
 */
type PaymentGwResult struct {
	RetCode     string  `json:"ret_code"`
	RetMsg      string  `json:"ret_msg"`
	OidPartner  string  `json:"oid_partner"`
	UserID      string  `json:"user_id"`
	TotalAmount float64 `json:"total_amount"`
	TxnSeqNo    string  `json:"txn_seqno"`
	AccpTxNo    string  `json:"accp_txno"`
	// 支付授权令牌，有效期30分钟。当交易需要二次验证的时候，需要通过token调用调用交易二次短信验证接口
	Token string `json:"token"`
	/*
	   网关地址。网银支付方式适用，返回跳转网关地址，用户跳转到网关完成后续支付操作。跳转方式：商户前端form表单POST提交。
	   扫码支付返回此地址后，按链接生成二维码，让用户扫码完成支付。
	*/
	GatewayURL string `json:"gateway_url"`
	// 支付参数集合。返回外部渠道的标准支付提交参数，微信/支付宝可参考官方文档。
	Payload string `json:"payload"`
}
