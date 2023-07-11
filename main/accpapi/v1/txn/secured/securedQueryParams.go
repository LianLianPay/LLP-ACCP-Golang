package secured

/**
 * 担保交易信息查询 请求参数
 */
type SecuredQueryParams struct {
	Timestamp  string `json:"timestamp"`
	OidPartner string `json:"oid_partner"`
	TxnSeqNo   string `json:"txn_seqno"`
	AccpTxNo   string `json:"accp_txno"`
}
