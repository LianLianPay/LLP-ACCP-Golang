package txn

/**
 * 银行卡快捷支付 响应参数
 */
type PaymentBankCardResult struct {
	LinkedAgrtno   string `json:"linked_agrtno"`
	AccountingDate string `json:"accounting_date"`
	FinishTime     string `json:"finish_time"`
	//golang中没有继承，我将main/accpapi/v1/txn/paymentResult.go中的参数拷贝进来
	RetCode     string  `json:"ret_code"`
	RetMsg      string  `json:"ret_msg"`
	OidPartner  string  `json:"oid_partner"`
	UserID      string  `json:"user_id"`
	Token       string  `json:"token"`
	TotalAmount float64 `json:"total_amount"`
	TxnSeqNo    string  `json:"txn_seqno"`
	AccpTxNo    string  `json:"accp_txno"`
}
