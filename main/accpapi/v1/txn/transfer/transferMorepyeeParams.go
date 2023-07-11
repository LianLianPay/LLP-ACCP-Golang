package transfer

/**
 * 内部代发申请 请求参数
 */
type TransferMorepyeeParams struct {
	Timestamp  string `json:"timestamp"`
	OidPartner string `json:"oid_partner"`
	NotifyURL  string `json:"notify_url"`
	/*
	   垫资标识。标识该笔代发交易是否支持平台商户垫资，适用于代发付款方为用户的业务场景。
	   默认：N
	   Y：支持垫资
	   N：不支持垫资
	*/
	FundsFlag string `json:"funds_flag"`
	/*
	   定向支付标识。
	   标识该笔代发交易是否是约定收款方的定向支付，默认：N。
	   Y：定向支付
	   N：普通支付
	*/
	DirectionalPayFlag string `json:"directionalpay_flag"`
	/*
	   连续代发标识。
	   标识该笔代发交易是否是连续代发，默认：N
	   Y：连续代发
	   N：普通代发
	   默认为N，如平台开通连续代发权限时，可以传入Y。
	   有效期30分钟，连续代发期间免密免验，超过需要再次验证。
	*/
	ContinuouslyFlag string `json:"continuously_flag"`
	RiskItem         string `json:"risk_item"`
	// 商户订单信息
	OrderInfo TransferMorepyeeOrderInfo `json:"order_info"`
	// 付款方信息
	PayerInfo TransferMorepyeePayerInfo `json:"payer_info"`
	// 收款方信息
	PayeeInfo []TransferMorepyeePayeeInfo `json:"payee_info"`
}
