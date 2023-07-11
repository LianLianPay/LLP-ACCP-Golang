package individual

/**
 * 个人用户开户申请 响应参数
 */
type OpenacctApplyResult struct {
	RetCode    string `json:"ret_code"`
	RetMsg     string `json:"ret_msg"`
	OidPartner string `json:"oid_partner"`
	UserID     string `json:"user_id"`
	UserName   string `json:"user_name"`
	TxnSeqno   string `json:"txn_seqno"`
	AccpTxno   string `json:"accp_txno"`
	Token      string `json:"token"`
}
