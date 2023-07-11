package refund

type RefundOrderInfo struct {
	// 退款订单号。标识一次退款请求，商户系统需要保证唯一
	RefundSeqNo string `json:"refund_seqno"`
	// 订单信息。用于订单说明，在异步通知透传返回
	OrderInfo string `json:"order_info"`
	// 退款订单时间。退款交易商户系统交易时间，格式:yyyyMMddHHmmss
	RefundTime string `json:"refund_time"`
	// 退款总金额
	RefundAmount float64 `json:"refund_amount"`
}
