package unlinked

type UnLinkedAcctIndApplyParams struct {
	Timestamp  string `json:"timestamp"`
	OidPartner string `json:"oid_partner"`
	UserID     string `json:"user_id"`
	TxnSeqno   string `json:"txn_seqno"`
	TxnTime    string `json:"txn_time"`
	NotifyURL  string `json:"notify_url"`
	// 绑定银行账号。个人用户绑定的银行卡号
	LinkedAcctno string `json:"linked_acctno"`
	// 支付密码
	Password  string `json:"password"`
	RandomKey string `json:"random_key"`
}
