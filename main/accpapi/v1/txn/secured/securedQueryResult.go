package secured

/**
 * 担保交易信息查询 响应参数
 */
type SecuredQueryResult struct {
	RetCode    string `json:"ret_code"`
	RetMsg     string `json:"ret_msg"`
	OidPartner string `json:"oid_partner"`
	AccpTxNo   string `json:"accp_txno"`
	// 商户订单信息
	OrderInfo SecuredQueryOrderInfo `json:"order_info"`
	// 付款方信息
	PayerInfo []SecuredQueryPayerInfo `json:"payer_info"`
	// 收款方信息
	PayeeInfo []SecuredQueryPayeeInfo `json:"payee_info"`
	// 担保支付订单
	PayInfo SecuredQueryPayInfo `json:"pay_info"`
	// 担保确认订单
	AcceptInfo []SecuredQueryAcceptInfo `json:"accept_info"`
	// 担保退款订单
	RefundInfo []SecuredQueryRefundInfo `json:"refund_info"`
}
