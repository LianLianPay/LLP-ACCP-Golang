package payment

type QueryPaymentOrderInfo struct {
	TxnSeqNo string `json:"txn_seqno"`
	// 商户系统交易时间
	TxnTime string `json:"txn_time"`
	// 订单总金额，单位为元
	TotalAmount float64 `json:"total_amount"`
	OrderInfo   string  `json:"order_info"`
}
