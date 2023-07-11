package enterprise

/**
 * 企业用户开户申请 请求参数
 */
type OpenacctApplyParams struct {
	Timestamp   string `json:"timestamp"`
	OidPartner  string `json:"oid_partner"`
	UserID      string `json:"user_id"`
	TxnSeqno    string `json:"txn_seqno"`
	TxnTime     string `json:"txn_time"`
	NotifyURL   string `json:"notify_url"`
	OpenSMSFlag string `json:"open_sms_flag"`
	RiskItem    string `json:"risk_item"`
	// 开户基本信息
	BasicInfo OpenacctApplyBasicInfo `json:"basicInfo"`
	// 开户绑卡信息
	LinkedAcctInfo OpenacctApplyLinkedAcctInfo `json:"linkedAcctInfo"`
	// 企业法定代表人信息
	LegalreptInfo OpenacctApplyLegalreptInfo `json:"legalreptInfo"`
	// 企业联系人信息
	ContactsInfo OpenacctApplyContactsInfo `json:"contactsInfo"`
	// 企业经营信息
	BusinessInfo OpenacctApplyBusinessInfo `json:"businessInfo"`
	// 企业基本户信息
	BasicAcctInfo OpenacctApplyBasicAcctInfo `json:"basicAcctInfo"`
	// 受益所有人信息uboInfos(可选，若传uboInfos，则按下列参数要求传)
	UboInfos OpenacctApplyUboInfos `json:"uboInfos"`
	// 开户账户申请信息
	ApplyAccountInfo OpenacctApplyAccountInfo `json:"applyAccountInfo"`
}
