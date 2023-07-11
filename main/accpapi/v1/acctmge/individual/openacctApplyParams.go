package individual

/**
 * 个人用户开户申请 请求参数
 */
type OpenacctApplyParams struct {
	Timestamp      string                      `json:"timestamp"`
	OidPartner     string                      `json:"oid_partner"`
	UserID         string                      `json:"user_id"`
	TxnSeqno       string                      `json:"txn_seqno"`
	TxnTime        string                      `json:"txn_time"`
	NotifyURL      string                      `json:"notify_url"`
	OpenSMSFlag    string                      `json:"open_sms_flag"`
	RiskItem       string                      `json:"risk_item"`
	BasicInfo      OpenacctApplyBasicInfo      `json:"basic_info"`
	LinkedAcctInfo OpenacctApplyLinkedAcctInfo `json:"linked_acct_info"`
	AccountInfo    OpenacctApplyAccountInfo    `json:"account_info"`
}
