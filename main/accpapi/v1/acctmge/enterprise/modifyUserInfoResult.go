package enterprise

/**
 * 企业用户信息修改 响应参数
 */
type ModifyUserInfoResult struct {
	RetCode    string `json:"ret_code"`
	RetMsg     string `json:"ret_msg"`
	OidPartner string `json:"oid_partner"`
	UserID     string `json:"user_id"`
	TxnSeqno   string `json:"txn_seqno"`
	AccpTxno   string `json:"accp_txno"`
	Remark     string `json:"remark"`
}
