package secured

type SecuredQueryPayeeInfo struct {
	/*
	   收款方类型。
	   用户：USER
	   平台商户：MERCHANT
	*/
	PayeeType string `json:"payee_type"`
	PayeeID   string `json:"payee_id"`
	// 收款金额
	Amount float64 `json:"amount"`
	// 已确认金额
	AcceptAmount float64 `json:"accept_amount"`
	// 退款金额
	RefundAmount float64 `json:"refund_amount"`
}
