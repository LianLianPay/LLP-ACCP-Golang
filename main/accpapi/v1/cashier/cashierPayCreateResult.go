package cashier

/**
 * 账户+收银台 响应参数
 */
type CashierPayCreateResult struct {
	RetCode     string  `json:"ret_code"`
	RetMsg      string  `json:"ret_msg"`
	OidPartner  string  `json:"oid_partner"`
	UserID      string  `json:"user_id"`
	TotalAmount float64 `json:"total_amount"`
	TxnSeqno    string  `json:"txn_seqno"`
	AccpTxno    string  `json:"accp_txno"`
	GatewayURL  string  `json:"gateway_url"`
}
