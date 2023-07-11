package find

/**
 * 找回密码申请 响应参数
 */
type FindPasswordApplyResult struct {
	RetCode    string `json:"ret_code"`
	RetMsg     string `json:"ret_msg"`
	OidPartner string `json:"oid_partner"`
	UserID     string `json:"user_id"`
	Token      string `json:"token"`
	RegPhone   string `json:"reg_phone"`
}
