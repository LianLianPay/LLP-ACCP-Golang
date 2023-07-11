package find

/**
 * 找回密码申请 请求参数
 */
type FindPasswordApplyParams struct {
	Timestamp    string `json:"timestamp"`
	OidPartner   string `json:"oid_partner"`
	UserID       string `json:"user_id"`
	LinkedAcctNo string `json:"linked_acctno"`
	RiskItem     string `json:"risk_item"`
}
