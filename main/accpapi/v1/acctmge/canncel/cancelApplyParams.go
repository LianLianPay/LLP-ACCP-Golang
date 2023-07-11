package canncel

/**
 * 销户申请 请求参数
 */
type CancelApplyParams struct {
	Timestamp  string `json:"timestamp"`
	OidPartner string `json:"oid_partner"`
	UserID     string `json:"user_id"`
	TxnSeqno   string `json:"txn_seqno"`
	TxnTime    string `json:"txn_time"`
	NotifyURL  string `json:"notify_url"`
	// 用户名称。企业用户传企业名称，个体工商户传营业执照的名称（营业执照上名称是****或者无的，传经营者姓名），个人用户传姓名。
	UserName string `json:"user_name"`
	/*
	   证件类型。
	   身份证：ID_CARD
	   统一社会信用代码证：UNIFIED_SOCIAL_CREDIT_CODE。
	*/
	IdType string `json:"id_type"`
	// 证件号码。 对应id_type的相关证件号码。
	IdNo string `json:"id_no"`
	// 绑定手机号。用户开户注册绑定手机号。
	RegPhone  string `json:"reg_phone"`
	Password  string `json:"password"`
	RandomKey string `json:"random_key"`
}
