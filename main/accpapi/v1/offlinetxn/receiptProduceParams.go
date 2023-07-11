package offlinetxn

/**
 * 电子回单生成 请求参数
 */
type ReceiptProduceParams struct {
	Timestamp  string `json:"timestamp"`
	OidPartner string `json:"oid_partner"`
	TxnSeqno   string `json:"txn_seqno"`
	TxnTime    string `json:"txn_time"`
	// 原交易流水号。二选一
	TradeTxnSeqno string `json:"trade_txn_seqno,omitempty"`
	// 原交易ACCP系统单号。二选一
	TradeAccpTxno string `json:"trade_accp_txno,omitempty"`
	/*
	   原交易订单类型。
	   PAYBILL:支付（包含充值，消费，内部代发，担保消费，担保确认）
	   CASHOUT:外部代发，提现
	   REFUND:退款
	*/
	TradeBillType string `json:"trade_bill_type"`
	// 订单总金额，单位为元，精确到小数点后两位。
	TotalAmount float64 `json:"total_amount"`
}
