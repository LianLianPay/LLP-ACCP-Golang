package offlinetxn

/**
 * 电子回单生成 响应参数
 */
type ReceiptProduceResult struct {
	RetCode       string `json:"ret_code"`
	RetMsg        string `json:"ret_msg"`
	OidPartner    string `json:"oid_partner"`
	TxnSeqno      string `json:"txn_seqno"`
	TradeTxnSeqno string `json:"trade_txn_seqno"`
	TradeAccpTxno string `json:"trade_accp_txno"`
	TotalAmount   string `json:"total_amount"`
	// 电子回单流水号。ACCP系统唯一定位一笔单子的电子回单。
	ReceiptAccpTxno string `json:"receipt_accp_txno"`
	// 授权令牌。有效期为120分钟。在电子回单下载接口需要传入。
	Token string `json:"token"`
}
