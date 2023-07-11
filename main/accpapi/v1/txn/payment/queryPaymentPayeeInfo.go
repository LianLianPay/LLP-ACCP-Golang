package payment

type QueryPaymentPayeeInfo struct {
	/*
	   收款方类型。
	   用户：USER
	   平台商户：MERCHANT
	*/
	PayeeType string `json:"payee_type"`
	// 收款方标识
	PayeeID string `json:"payee_id"`
	// 付款金额
	Amount string `json:"amount"`
}
