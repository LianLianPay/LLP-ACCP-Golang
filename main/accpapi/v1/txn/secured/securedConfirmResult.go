package secured

/**
 * 担保交易确认 响应参数
 */
type SecuredConfirmResult struct {
	RetCode         string  `json:"ret_code"`
	RetMsg          string  `json:"ret_msg"`
	OidPartner      string  `json:"oid_partner"`
	UserID          string  `json:"user_id"`
	TxnSeqNo        string  `json:"txn_seqno"`
	TotalAmount     float64 `json:"total_amount"`
	AccpTxNo        string  `json:"accp_txno"`
	AccpConfirmTxNo string  `json:"accp_confirm_txno"`
}
