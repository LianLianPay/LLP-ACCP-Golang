package txn

/**
 * 交易二次短信验证 请求参数
 */
type ValidationSmsParams struct {
	Timestamp   string `json:"timestamp"`
	OidPartner  string `json:"oid_partner"`
	PayerType   string `json:"payer_type"`
	PayerID     string `json:"payer_id"`
	TxnSeqNo    string `json:"txn_seqno"`
	TotalAmount string `json:"total_amount"`
	Token       string `json:"token"`
	VerifyCode  string `json:"verify_code"`
}
