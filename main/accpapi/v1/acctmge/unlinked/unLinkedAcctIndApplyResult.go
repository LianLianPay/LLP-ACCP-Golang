package unlinked

/**
 * 个人用户解绑银行卡 响应参数
 */
type UnLinkedAcctIndApplyResult struct {
	RetCode    string `json:"ret_code"`
	RetMsg     string `json:"ret_msg"`
	OidPartner string `json:"oid_partner"`
	UserID     string `json:"user_id"`
	TxnSeqno   string `json:"txn_seqno"`
	AccpTxno   string `json:"accp_txno"`
}
