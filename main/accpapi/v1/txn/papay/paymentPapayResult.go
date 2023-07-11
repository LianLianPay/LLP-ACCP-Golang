package papay

/**
 * 委托代扣 响应参数
 */
type PaymentPapayResult struct {
	RetCode     string  `json:"ret_code"`
	RetMsg      string  `json:"ret_msg"`
	OidPartner  string  `json:"oid_partner"`
	UserID      string  `json:"user_id"`
	TxnSeqNo    string  `json:"txn_seqno"`
	TotalAmount float64 `json:"total_amount"`
	AccpTxNo    string  `json:"accp_txno"`
	/*
	   支付交易状态。
	   TRADE_WAIT_PAY:交易处理中
	   TRADE_SUCCESS:交易成功
	   TRADE_CLOSE:交易失败
	   支付交易状态以此为准，商户必须依据此字段值进行后续业务逻辑处理。
	*/
	TxnStatus      string `json:"txn_status"`
	AccountingDate string `json:"accounting_date"`
	FinishTime     string `json:"finish_time"`
	// 渠道失败原因。渠道返回失败时,返回具体失败原因。
	FailReason string `json:"fail_reason"`
}
