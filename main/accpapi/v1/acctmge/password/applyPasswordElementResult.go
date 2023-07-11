package password

/**
 * 申请密码控件Token 响应参数
 */
type ApplyPasswordElementResult struct {
	RetCode    string `json:"ret_code"`
	RetMsg     string `json:"ret_msg"`
	OidPartner string `json:"oid_partner"`
	// 用于唤起密码控件的token
	PasswordElementToken string `json:"password_element_token"`
}
