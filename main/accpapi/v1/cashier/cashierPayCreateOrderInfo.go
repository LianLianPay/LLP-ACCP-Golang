package cashier

type CashierPayCreateOrderInfo struct {
	TxnSeqno    string  `json:"txn_seqno"`
	TxnTime     string  `json:"txn_time"`
	TotalAmount float64 `json:"total_amount"`
	OrderInfo   string  `json:"order_info"`
	GoodsName   string  `json:"goods_name"`
	GoodsURL    string  `json:"goods_url"`
}
