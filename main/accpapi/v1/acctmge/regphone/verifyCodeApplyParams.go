package regphone

/**
 * 绑定手机验证码申请 请求参数
 */
type VerifyCodeApplyParams struct {
	Timestamp  string `json:"timestamp"`
	OidPartner string `json:"oid_partner"`
	UserID     string `json:"user_id"`
	RegPhone   string `json:"reg_phone"`
}
