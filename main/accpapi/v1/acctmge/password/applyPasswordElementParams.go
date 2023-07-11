package password

type ApplyPasswordElementParams struct {
	Timestamp  string `json:"timestamp"`
	OidPartner string `json:"oid_partner"`
	// 商户名称，余额支付场景必填
	PartnerName string `json:"partner_name"`
	// 用户在商户系统中的唯一标识，修改密码场景必填
	UserID   string `json:"user_id"`
	TxnSeqno string `json:"txn_seqno"`
	// 交易金额，单位：元，精确到小数点后两位，例如:1.00，支付和提现场景必填
	Amount float64 `json:"amount"`
	/*
	   密码使用场景：
	   设置密码：setting_password
	   修改密码：change_password
	   换绑卡：bind_card_password
	   提现密码：cashout_password
	   支付密码：pay_password
	*/
	PasswordScene    string `json:"password_scene"`
	EncryptAlgorithm string `json:"encrypt_algorithm"`
	FlagChnl         string `json:"flag_chnl"`
	// 收款人姓名，提现场景必填。
	PyeeName string `json:"pyee_name"`
}
