package papay

type PaymentPapayPayerInfo struct {
	/*
	   付款方类型。
	   用户：USER
	   平台商户：MERCHANT
	*/
	PayerType string `json:"payer_type"`
	// 付款方标识。付款方为用户时设置user_id。
	PayerID string `json:"payer_id"`
	// 委托代扣协议id
	ContractID string `json:"contract_id"`
}
