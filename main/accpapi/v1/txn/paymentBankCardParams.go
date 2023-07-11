package txn

/**
 * 银行卡快捷支付 请求参数
 */
type PaymentBankCardParams struct {
	Timestamp          string  `json:"timestamp"`
	OidPartner         string  `json:"oid_partner"`
	TxnSeqNo           string  `json:"txn_seqno"`
	TotalAmount        float64 `json:"total_amount"`
	RiskItem           string  `json:"risk_item"`
	DirectionalPayFlag string  `json:"directionalpay_flag"`
	// 付款方信息
	PayerInfo PaymentPayerInfo `json:"payerInfo"`
	// 付款银行卡信息
	BankCardInfo PaymentBankCardInfo `json:"bankCardInfo"`
	// 付款方式信息
	PayMethods []PaymentBankCardPayMethods `json:"payMethods"`
}
