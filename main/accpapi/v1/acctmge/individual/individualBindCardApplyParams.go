package individual

/**
 * 个人用户新增绑卡 请求参数
 */
type IndividualBindCardApplyParams struct {
	Timestamp    string `json:"timestamp"`
	OidPartner   string `json:"oid_partner"`
	UserID       string `json:"user_id"`
	TxnSeqno     string `json:"txn_seqno"`
	TxnTime      string `json:"txn_time"`
	NotifyURL    string `json:"notify_url"`
	RiskItem     string `json:"risk_item"`
	LinkedAcctNo string `json:"linked_acctno"`
	LinkedPhone  string `json:"linked_phone"`
	Password     string `json:"password"`
	RandomKey    string `json:"random_key"`
}
