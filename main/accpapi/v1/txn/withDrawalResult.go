package txn

/**
 * 提现申请 响应参数
 */
type WithDrawalResult struct {
	RetCode     string  `json:"ret_code"`
	RetMsg      string  `json:"ret_msg"`
	OidPartner  string  `json:"oid_partner"`
	UserID      string  `json:"user_id"`
	TxnSeqNo    string  `json:"txn_seqno"`
	TotalAmount float64 `json:"total_amount"`
	FeeAmount   float64 `json:"fee_amount"`
	AccpTxNo    string  `json:"accp_txno"`
	Token       string  `json:"token"`
}
