package txn

/**
 * 网关类支付 请求参数
 */
type PaymentGwParams struct {
	Timestamp   string  `json:"timestamp"`
	OidPartner  string  `json:"oid_partner"`
	TxnSeqno    string  `json:"txn_seqno"`
	TotalAmount float64 `json:"total_amount"`
	RiskItem    string  `json:"risk_item"`
	// 应用ID。开发者在微信或支付宝开放平台申请的APPID。非网银类，除扫码支付，都必传
	AppID string `json:"appid"`
	// 渠道用户标识。微信公众号和微信小程序支付时必传。此参数为微信用户在商户对应APPID下的唯一标识或支付宝买家的支付宝唯一用户号
	OpenID string `json:"openid"`
	// 银行编码。付款方式为网银类时可指定。
	Bankcode string `json:"bankcode"`
	// 设备标识。自定义参数，可以为终端设备号。
	DeviceInfo string `json:"device_info"`
	// 终端用户IP。支持IPV4和IPV6两种格式的IP地址。
	ClientIP      string `json:"client_ip"`
	DeviceVersion string `json:"device_version"`
	/*
	   限制卡类型。限制某种卡类型支付权限，支付宝和微信支付都适用，若限制多种类型以“,”分隔（暂时只支持限制信用卡支付）。credit：限制适用信用卡支付。
	*/
	LimitCardType string `json:"limit_card_type"`
	/*
	   业务扩展字段。渠道扩展字段JSON串，若渠道为支付宝，则支持的键值对如下：
	   sys_service_provider_id
	   hb_fq_num
	   hb_fq_seller_percent
	   industry_reflux_info
	   card_type
	   支付宝微信同时支持键值：accp_sub_mch_id 子商户号
	   微信H5扩展参数：req_domain 来源域名
	*/
	ExtendParams string `json:"extend_params"`
	/*
	   定向支付标识。
	   标识该笔交易是否是约定收款方的定向支付，默认：N。
	   Y：定向支付
	   N：普通支付。
	   组合余额支付场景时适用。
	*/
	DirectionalPayFlag string `json:"directionalpay_flag"`
	// 付款方信息
	PayerInfo PaymentPayerInfo `json:"payer_info"`
	// 付款方式信息
	PayMethods []PaymentBankCardPayMethods `json:"payMethods"`
}
