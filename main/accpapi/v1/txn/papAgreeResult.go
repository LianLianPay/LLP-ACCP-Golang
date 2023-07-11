package txn

/**
 * 用户委托协议申请 响应参数
 */
type PapAgreeResult struct {
	RetCode    string `json:"ret_code"`
	RetMsg     string `json:"ret_msg"`
	GatewayURL string `json:"gateway_url"`
	OidPartner string `json:"oid_partner"`
	UserID     string `json:"user_id"`
	TxnSeqNo   string `json:"txn_seqno"`
	AccpTxNo   string `json:"accp_txno"`
}
