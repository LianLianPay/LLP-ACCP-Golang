package individual

/**
 * 个人用户信息修改 请求参数
 */
type IndividualModifyUserInfoParams struct {
	Timestamp  string                            `json:"timestamp"`
	OidPartner string                            `json:"oid_partner"`
	UserID     string                            `json:"user_id"`
	TxnSeqno   string                            `json:"txn_seqno"`
	TxnTime    string                            `json:"txn_time"`
	NotifyURL  string                            `json:"notify_url"`
	BasicInfo  IndividualModifyUserInfoBasicInfo `json:"basicInfo"`
}
