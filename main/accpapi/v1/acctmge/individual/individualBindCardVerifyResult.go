package individual

/**
 * 个人用户新增绑卡验证 响应参数
 */
type IndividualBindCardVerifyResult struct {
	RetCode      string `json:"ret_code"`
	RetMsg       string `json:"ret_msg"`
	OidPartner   string `json:"oid_partner"`
	UserID       string `json:"user_id"`
	TxnSeqno     string `json:"txn_seqno"`
	AccpTxno     string `json:"accp_txno"`
	LinkedAgrtno string `json:"linked_agrtno"`
}
