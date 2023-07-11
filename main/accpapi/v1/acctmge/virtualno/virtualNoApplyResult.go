package virtualno

/**
 * 虚拟卡申请 请求参数
 */
type VirtualNoApplyResult struct {
	RetCode    string `json:"ret_code"`
	RetMsg     string `json:"ret_msg"`
	OidPartner string `json:"oid_partner"`
	UserID     string `json:"user_id"`
	TxnSeqno   string `json:"txn_seqno"`
	// 虚拟卡卡号，转账支付关联的虚拟卡。
	VirtualNo string `json:"virtualno"`
	// 申请人姓名。申请虚拟卡的客户姓名。
	ApplyUsername string `json:"apply_username"`
	// 申请人ID
	ApplyUserID string `json:"apply_userid"`
	/*
	   申请用途。
	   消费：consume
	   现金上缴：deposit_cash
	   费用上缴：deposit_fee
	   其他：other
	*/
	ApplySence string `json:"apply_sence"`
	/*
	   虚拟卡业务类型。
	   消费业务：OFFLINE_CONSUME
	*/
	BusinessType string `json:"bussiness_type"`
	// 资金挂账通知地址
	BBocNotify string `json:"bboc_notify"`
}
