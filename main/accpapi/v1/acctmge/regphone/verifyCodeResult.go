package regphone

/**
 * 绑定手机验证码申请/验证 响应参数
 */
type VerifyCodeResult struct {
	RetCode    string `json:"ret_code"`
	RetMsg     string `json:"ret_msg"`
	OidPartner string `json:"oid_partner"`
	UserID     string `json:"user_id"`
	RegPhone   string `json:"reg_phone"`
}
