package find

/**
 * 找回密码验证 响应参数
 */
type FindPasswordVerifyResult struct {
	RetCode    string `json:"ret_code"`
	RetMsg     string `json:"ret_msg"`
	OidPartner string `json:"oid_partner"`
	UserID     string `json:"user_id"`
}
