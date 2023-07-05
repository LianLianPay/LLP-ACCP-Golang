package txn

type TradeCreateOrderInfo struct {
	TxnSeqno    string
	TxnTime     string
	TotalAmount float64
	FeeAmount   float64
	OrderInfo   string
	GoodsName   string
	GoodsURL    string
}
