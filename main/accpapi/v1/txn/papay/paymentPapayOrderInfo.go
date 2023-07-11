package papay

type PaymentPapayOrderInfo struct {
	TxnSeqno    string  `json:"txn_seqno"`
	TxnTime     string  `json:"txn_time"`
	TotalAmount float64 `json:"total_amount"`
	OrderInfo   string  `json:"order_info"`
	GoodsName   string  `json:"goods_name"`
	GoodsURL    string  `json:"goods_url"`
	/*
	   优惠券支付模式：默认，CONSUME_SUCCESS_PAY。
	   CONSUME_SUCCESS_PAY：支付成功后，转到担保账户。
	   SECURED_CONFIRM_AUTO：担保确认后转到担保账户，优惠券金额由系统计算。
	   SECURED_CONFIRM_SUBMIT：担保确认后转到担保账户，优惠券金额由商户上送。
	*/
	CouponPayMode string `json:"coupon_pay_mode"`
}
