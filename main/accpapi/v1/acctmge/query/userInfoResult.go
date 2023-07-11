package query

type UserInfoResult struct {
	RetCode    string `json:"ret_code"`
	RetMsg     string `json:"ret_msg"`
	OidPartner string `json:"oid_partner"`
	UserID     string `json:"user_id"`
	UserName   string `json:"user_name"`
	OidUserNo  string `json:"oid_userno"`
	/*
	   用户状态。
	   ACTIVATE_PENDING :已登记或开户失败（原待激活）
	   CHECK_PENDING :审核中（原待审核）
	   REMITTANCE_VALID_PENDING :审核通过，待打款验证（企业用户使用，暂未要求）
	   NORMAL :正常
	   CANCEL :销户
	   PAUSE :暂停
	   ACTIVATE_PENDING_NEW ：待激活
	*/
	UserStatus   string `json:"user_status"`
	BankOpenFlag string `json:"bank_open_flag"`
	Remark       string `json:"remark"`
	BankAccount  string `json:"bank_account"`
}
