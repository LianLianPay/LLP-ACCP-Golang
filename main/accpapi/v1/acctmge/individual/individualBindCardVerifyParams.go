package individual

/**
 * 个人用户新增绑卡验证 请求参数
 */
type IndividualBindCardVerifyParams struct {
	Timestamp  string `json:"timestamp"`
	OidPartner string `json:"oid_partner"`
	UserID     string `json:"user_id"`
	TxnSeqno   string `json:"txn_seqno"`
	Token      string `json:"token"`
	VerifyCode string `json:"verify_code"`
}
