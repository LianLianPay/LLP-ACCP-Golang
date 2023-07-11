package offlinetxn

/**
 * 电子回单下载 请求参数
 */
type ReceiptDownloadParams struct {
	Timestamp  string `json:"timestamp"`
	OidPartner string `json:"oid_partner"`
	// 授权令牌，有效期为120分钟。
	Token string `json:"token"`
	// 电子回单流水号。
	ReceiptAccpTxno string `json:"receipt_accp_txno"`
}
