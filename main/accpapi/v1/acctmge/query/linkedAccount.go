package query

type LinkedAccount struct {
	// 个人用户绑定的银行卡号，企业用户绑定的银行帐号。
	LinkedAcctNo string `json:"linked_acctno"`
	// 银行编码
	LinkedBankCode string `json:"linked_bankcode"`
	// 企业用户开户行行号
	LinkedBrBankNo string `json:"linked_brbankno"`
	// 对公账户开户行名。企业用户绑定账户开户支行名称。
	LinkedBrBankName string `json:"linked_brbankname"`
	// 银行预留手机号
	LinkedPhone string `json:"linked_phone"`
	// 绑卡协议号
	LinkedAgrtNo string `json:"linked_agrtno"`
}
