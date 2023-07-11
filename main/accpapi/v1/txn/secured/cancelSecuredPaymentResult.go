package secured

/**
 * 担保确认撤销 响应参数
 */
type CancelSecuredPaymentResult struct {
	RetCode     string  `json:"ret_code"`
	RetMsg      string  `json:"ret_msg"`
	OidPartner  string  `json:"oid_partner"`
	AccpTxNo    string  `json:"accp_txno"`
	TxnSeqNo    string  `json:"txn_seqno"`
	TotalAmount float64 `json:"total_amount"`
	// 收款方信息
	PayeeInfo []CancelSecuredPaymentPayeeInfo `json:"payee_info"`
}
