package virtualno

/**
 * 虚拟卡申请 请求参数
 */
type VirtualNoApplyParams struct {
	Timestamp  string `json:"timestamp"`
	OidPartner string `json:"oid_partner"`
	TxnSeqno   string `json:"txn_seqno"`
	TxnTime    string `json:"txn_time"`
	// 虚拟卡参数
	VirtualInfo VirtualInfo `json:"virtual_info"`
}
