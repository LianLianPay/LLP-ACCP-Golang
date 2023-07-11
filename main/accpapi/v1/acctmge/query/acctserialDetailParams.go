package query

/**
 * 资金流水详情查询 请求参数
 */
type AcctserialDetailParams struct {
	Timestamp  string `json:"timestamp"`
	OidPartner string `json:"oid_partner"`
	UserID     string `json:"user_id"`
	UserType   string `json:"user_type"`
	JnoAcct    string `json:"jno_acct"`
}
