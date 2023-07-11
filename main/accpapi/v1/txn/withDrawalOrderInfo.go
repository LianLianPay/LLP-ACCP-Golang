package txn

type WithDrawalOrderInfo struct {
	TxnSeqNo    string  `json:"txn_seqno"`
	TxnTime     string  `json:"txn_time"`
	TotalAmount float64 `json:"total_amount"`
	FeeAmount   float64 `json:"fee_amount"`
	OrderInfo   string  `json:"order_info"`
	// 交易附言。提现交易附言。单笔金额大于等于5w必须提供
	Postscript string `json:"postscript"`
}
