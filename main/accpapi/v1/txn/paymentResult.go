package txn

/**
 * 余额支付/银行卡快捷支付（Base类） 响应参数
 */
type PaymentResult struct {
	RetCode     string  `json:"ret_code"`
	RetMsg      string  `json:"ret_msg"`
	OidPartner  string  `json:"oid_partner"`
	UserID      string  `json:"user_id"`
	Token       string  `json:"token"`
	TotalAmount float64 `json:"total_amount"`
	TxnSeqNo    string  `json:"txn_seqno"`
	AccpTxNo    string  `json:"accp_txno"`
}
