package acctmgr

/**
 * 匿名用户解绑银行卡 请求参数
 */
type UnLinkedAcctIndApplyParams struct {
	Timestamp  string `json:"timestamp"`
	OidPartner string `json:"oid_partner"`
	TxnSeqNo   string `json:"txn_seqno"`
	UserID     string `json:"user_id"`
	TxnTime    string `json:"txn_time"`
	NotifyURL  string `json:"notify_url"`
	// 绑定银行账号。该字段需要RSA公钥加密传输
	LinkedAcctNo string `json:"linked_acctno"`
}
