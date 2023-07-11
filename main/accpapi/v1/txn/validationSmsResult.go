package txn

/**
 * 交易二次短信验证 响应参数
 */
type ValidationSmsResult struct {
	RetCode        string  `json:"ret_code"`
	RetMsg         string  `json:"ret_msg"`
	OidPartner     string  `json:"oid_partner"`
	UserID         string  `json:"user_id"`
	TxnSeqNo       string  `json:"txn_seqno"`
	TotalAmount    float64 `json:"total_amount"`
	AccpTxNo       string  `json:"accp_txno"`
	AccountingDate string  `json:"accounting_date"`
	FinishTime     string  `json:"finish_time"`
}
