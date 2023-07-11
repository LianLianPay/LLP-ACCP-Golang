package regphone

/**
 * 绑定手机验证码验证 请求参数
 */
type VerifyCodeVerifyParams struct {
	Timestamp  string `json:"timestamp"`
	OidPartner string `json:"oid_partner"`
	UserID     string `json:"user_id"`
	RegPhone   string `json:"reg_phone"`
	VerifyCode string `json:"verify_code"`
}
