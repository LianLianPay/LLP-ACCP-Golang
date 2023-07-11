package find

/**
 * 找回密码验证 请求参数
 */
type FindPasswordVerifyParams struct {
	Timestamp  string `json:"timestamp"`
	OidPartner string `json:"oid_partner"`
	UserID     string `json:"user_id"`
	Token      string `json:"token"`
	VerifyCode string `json:"verify_code"`
	RandomKey  string `json:"random_key"`
	Password   string `json:"password"`
}
