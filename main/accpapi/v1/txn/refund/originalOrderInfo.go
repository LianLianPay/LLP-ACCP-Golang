package refund

type OriginalOrderInfo struct {
	TxnSeqNo string `json:"txn_seqno"`
	// 订单总金额，单位为元，精确到小数点后两位
	TotalAmount float64 `json:"total_amount"`
}
