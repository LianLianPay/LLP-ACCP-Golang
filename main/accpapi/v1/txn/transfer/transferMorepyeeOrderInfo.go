package transfer

type TransferMorepyeeOrderInfo struct {
	TxnSeqNo    string  `json:"txn_seqno"`
	TxnTime     string  `json:"txn_time"`
	TotalAmount float64 `json:"total_amount"`
	/*
	   代发交易用途。
	   服务费
	   信息费
	   修理费
	   佣金支付
	   贷款
	   其他
	*/
	TxnPurpose string `json:"txn_purpose"`
	// 订单信息，在查询API和支付通知中原样返回，可作为自定义参数使用
	OrderInfo string `json:"order_info"`
}
