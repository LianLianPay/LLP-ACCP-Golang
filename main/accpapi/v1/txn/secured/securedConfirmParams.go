package secured

/**
 * 担保交易确认 请求参数
 */
type SecuredConfirmParams struct {
	Timestamp  string `json:"timestamp"`
	OidPartner string `json:"oid_partner"`
	UserID     string `json:"user_id"`
	/*
	   确认方式。
	   ALL：全订单金额一次性确认（创单时指定了收款方）
	   PART：订单金额部分多次确认
	*/
	ConfirmMode string `json:"confirm_mode"`
	// 原商户订单信息
	OriginalOrderInfo SecuredConfirmOriginalOrderInfo `json:"original_order_info"`
	// 确认订单信息
	ConfirmOrderInfo SecuredConfirmOrderInfo `json:"confirm_order_info"`
	// 确认收款方信息
	PayeeInfo []SecuredConfirmPayeeInfo `json:"payee_info"`
}
