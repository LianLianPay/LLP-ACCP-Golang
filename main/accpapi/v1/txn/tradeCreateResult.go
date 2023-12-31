package txn

/**
 * 支付统一创单 响应参数
 */
type TradeCreateResult struct {
	RetCode     string  `json:"ret_code"`
	RetMsg      string  `json:"ret_msg"`
	OidPartner  string  `json:"oid_partner"`
	UserID      string  `json:"user_id"`
	TotalAmount float64 `json:"total_amount"`
	TxnSeqNo    string  `json:"txn_seqno"`
	AccpTxNo    string  `json:"accp_txno"`
}
