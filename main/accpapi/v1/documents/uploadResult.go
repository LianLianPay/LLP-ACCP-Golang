package documents

/**
 * 文件上传 响应参数
 */
type UploadResult struct {
	RetCode    string `json:"ret_code"`
	RetMsg     string `json:"ret_msg"`
	OidPartner string `json:"oid_partner"`
	DocID      string `json:"doc_id"`
	TxnSeqno   string `json:"txn_seqno"`
}
