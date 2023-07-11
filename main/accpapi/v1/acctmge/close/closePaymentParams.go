package close

/**
 * 支付单关单 请求参数
 */
type ClosePaymentParams struct {
	Timestamp  string `json:"timestamp"`
	OidPartner string `json:"oid_partner"`
	// 商户系统唯一交易流水号。二选一
	TxnSeqno string `json:"txn_seqno,omitempty"`
	// ACCP系统交易单号。二选一
	AccpTxno string `json:"accp_txno,omitempty"`
}
