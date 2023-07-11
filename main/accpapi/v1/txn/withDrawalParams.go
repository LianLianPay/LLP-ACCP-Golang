package txn

/**
 * 提现申请 请求参数
 */
type WithDrawalParams struct {
	Timestamp  string `json:"timestamp"`
	OidPartner string `json:"oid_partner"`
	NotifyURL  string `json:"notify_url"`
	PayExpire  string `json:"pay_expire"`
	RiskItem   string `json:"risk_item"`
	// 绑定银行账号
	LinkedAcctNo string `json:"linked_acctno"`
	// 绑卡协议号
	LinkedAgrtNo string `json:"linked_agrtno"`
	/*
	   垫资标识。标识该笔提现交易是否支持平台商户垫资，适用于用户提现业务场景。
	   默认：N
	   Y：支持垫资
	   N：不支持垫资
	*/
	FundsFlag string `json:"funds_flag"`
	/*
	   审核标识。标识该笔订单是否需要审核，默认:N
	   Y:需要提现确认
	   N：无需提现确认
	*/
	CheckFlag string `json:"check_flag"`
	/*
	   到账类型。
	   TRANS_THIS_TIME :实时到账
	   TRANS_NORMAL :普通到账（2小时内）
	   TRANS_NEXT_TIME :次日到账
	   默认：实时到账。
	   商户订单信息orderInfo
	*/
	PayTimeType string `json:"pay_time_type"`
	// 商户订单信息
	OrderInfo WithDrawalOrderInfo `json:"orderInfo"`
	// 付款方信息
	PayerInfo WithDrawalPayerInfo `json:"payerInfo"`
}
