package close

/**
 * 支付单关单 响应参数
 */
type ClosePaymentResult struct {
	RetCode     string `json:"ret_code"`
	RetMsg      string `json:"ret_msg"`
	OidPartner  string `json:"oid_partner"`
	UserID      string `json:"user_id"`
	TxnSeqno    string `json:"txn_seqno"`
	AccpTxno    string `json:"accp_txno"`
	TotalAmount string `json:"total_amount"`
}
