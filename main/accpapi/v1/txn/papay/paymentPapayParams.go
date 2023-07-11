package papay

/**
 * 委托代扣 请求参数
 */
type PaymentPapayParams struct {
	Timestamp  string `json:"timestamp"`
	OidPartner string `json:"oid_partner"`
	/*
	   交易类型。
	   用户充值：USER_TOPUP
	   商户充值：MCH_TOPUP
	   普通消费：GENERAL_CONSUME
	   担保消费：SECURED_CONSUME
	*/
	TxnType string `json:"txn_type"`
	UserID  string `json:"user_id"`
	/*
	   用户类型。默认：注册用户。
	   注册用户：REGISTERED
	   匿名用户：ANONYMOUS
	*/
	UserType  string `json:"user_type"`
	NotifyURL string `json:"notify_url"`
	ReturnURL string `json:"return_url"`
	AppID     string `json:"appid"`
	ClientIP  string `json:"client_ip"`
	RiskItem  string `json:"risk_item"`
	// 商户订单信息
	OrderInfo PaymentPapayOrderInfo `json:"orderInfo"`
	// 付款方信息
	PayerInfo PaymentPapayPayerInfo `json:"payerInfo"`
	// 付款方式信息
	PayMethods PaymentPapayPayMethod `json:"payMethods"`
	// 收款方信息
	PayeeInfo PaymentPapayPayeeInfo `json:"payeeInfo"`
	// 微信附加字段
	WechatInfo WechatInfo `json:"wechatInfo"`
	// 支付宝附加字段
	AlipayInfo AlipayInfo `json:"alipayInfo"`
}
