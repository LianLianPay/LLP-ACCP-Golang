package refund

type PyeeRefundInfo struct {
	// 原收款方id，本次退款需要处理的原交易收款方id
	PayeeID string `json:"payee_id"`
	/*
	   原收款方类型。
	   用户：USER
	   平台商户：MERCHANT
	*/
	PayeeType string `json:"payee_type"`
	/*
	   原收款方账户类型。
	   用户账户：USEROWN
	   平台商户自有资金账户：MCHOWN
	   平台商户担保账户：MCHASSURE
	   平台商户优惠券账户：MCHCOUPON
	   平台商户手续费账户：MCHFEE
	*/
	PayeeAcctType string `json:"payee_accttype"`
	// 退款金额。本次需要退款的金额，不允许超过对应原收款方的收款金额
	PayeeRefundAmount float64 `json:"payee_refund_amount"`
	// 垫资标识。当原收款方金额不足时，是否由平台垫资的标识，默认:N
	IsAdvancePay string `json:"is_advance_pay"`
}
