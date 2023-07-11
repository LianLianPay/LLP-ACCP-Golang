package secured

type SecuredConfirmOriginalOrderInfo struct {
	TxnSeqNo    string  `json:"txn_seqno"`
	TotalAmount float64 `json:"total_amount"`
}
