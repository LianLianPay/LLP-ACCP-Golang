package payment

type QueryPaymentPayerInfo struct {
	/*
	   付款方类型。
	   用户：USER
	   平台商户：MERCHANT
	*/
	PayerType string `json:"payer_type"`
	/*
	   付款方标识。
	   付款方为用户时设置user_id
	   付款方为商户时设置平台商户号
	*/
	PayerID string `json:"payer_id"`
	// 付款方式
	Method string `json:"method"`
	// 付款金额。付款方式对应的金额
	Amount float64 `json:"amount"`
}
