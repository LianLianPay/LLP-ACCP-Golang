package webapge

/**
 * 用户开户申请(页面接入) 响应参数
 */
type OpenacctApplyResult struct {
	RetCode    string `json:"ret_code"`
	RetMsg     string `json:"ret_msg"`
	OidPartner string `json:"oid_partner"`
	UserID     string `json:"user_id"`
	TxnSeqno   string `json:"txn_seqno"`
	AccpTxno   string `json:"accp_txno"`
	GatewayURL string `json:"gateway_url"`
}
