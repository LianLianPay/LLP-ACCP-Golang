package enterprise

/**
 * 企业用户信息修改 请求参数
 */
type ModifyUserInfoParams struct {
	Timestamp  string `json:"timestamp"`
	OidPartner string `json:"oid_partner"`
	UserID     string `json:"user_id"`
	TxnSeqno   string `json:"txn_seqno"`
	TxnTime    string `json:"txn_time"`
	NotifyURL  string `json:"notify_url"`
	// 基本信息更新
	BasicInfo OpenacctApplyBasicInfo `json:"basicInfo"`
	// 企业法定代表人信息更新
	LegalreptInfo OpenacctApplyLegalreptInfo `json:"legalreptInfo"`
	// 企业联系人信息更新
	ContactsInfo OpenacctApplyContactsInfo `json:"contactsInfo"`
	// 企业经营信息更新
	BusinessInfo OpenacctApplyBusinessInfo `json:"businessInfo"`
	// 已有受益所有人信息替换
	UboInfos []OpenacctApplyUboInfos `json:"uboInfos"`
}
