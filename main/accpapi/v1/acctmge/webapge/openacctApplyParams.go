package webapge

/**
 * 用户开户申请(页面接入) 请求参数
 */
type OpenacctApplyParams struct {
	Timestamp           string `json:"timestamp"`
	OidPartner          string `json:"oid_partner"`
	UserID              string `json:"user_id"`
	TxnSeqno            string `json:"txn_seqno"`
	TxnTime             string `json:"txn_time"`
	FlagChnl            string `json:"flag_chnl"`
	ReturnURL           string `json:"return_url"`
	NotifyURL           string `json:"notify_url"`
	UserType            string `json:"user_type"`
	CustTradeSerialType string `json:"cust_trade_serial_type"`
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
	// 页面模板定制
	Theme OpenacctApplyTheme `json:"theme"`
	// 开户账户申请信息
	AccountInfo OpenacctApplyAccountInfo `json:"accountInfo"`
}
