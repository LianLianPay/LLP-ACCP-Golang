package secured

/**
 * 担保确认撤销 请求参数
 */
type CancelSecuredPaymentParams struct {
	Timestamp  string `json:"timestamp"`
	OidPartner string `json:"oid_partner"`
	// 商户担保确认单号。传担保确认接口传的confirm_seqno的值，与accp_txno二选一，建议优先使用ACCP系统交易单号。
	TxnSeqNo string `json:"txn_seqno"`
	// ACCP系统交易单号。传担保确认接口返回的accp_confirm_txno，与txn_seqno二选一，建议优先使用ACCP系统交易单号。
	AccpTxNo string `json:"accp_txno"`
}
