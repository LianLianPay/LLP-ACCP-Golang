package payment

/**
 * 支付结果查询 请求参数
 */
type QueryPaymentParams struct {
	Timestamp  string `json:"timestamp"`
	OidPartner string `json:"oid_partner"`
	// 商户系统唯一交易流水号。【三选一】，建议优先使用ACCP系统交易单号
	TxnSeqNo string `json:"txn_seqno"`
	// ACCP系统交易单号。【三选一】，建议优先使用ACCP系统交易单号
	AccpTxNo string `json:"accp_txno"`
	// 上游渠道流水号。【三选一】，建议优先使用 ACCP 系统交易单号
	SubChnlNo string `json:"sub_chnl_no"`
}
