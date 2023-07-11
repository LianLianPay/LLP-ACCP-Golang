package offlinetxn

/**
 * 电子回单下载 响应参数
 */
type ReceiptDownloadResult struct {
	RetCode         string  `json:"ret_code"`
	RetMsg          string  `json:"ret_msg"`
	OidPartner      string  `json:"oid_partner"`
	TradeTxnSeqno   string  `json:"trade_txn_seqno"`
	TradeAccpTxno   string  `json:"trade_accp_txno"`
	TotalAmount     float64 `json:"total_amount"`
	ReceiptAccpTxno string  `json:"receipt_accp_txno"`
	/*
	   电子回单生成状态。
	   SUCCESS:生成成功
	   HANDING:生成中
	   UNSUPPORT:不支持电子回单
	*/
	ReceiptStatus string `json:"receipt_status"`
	// 电子回单生集合文件。所有加密后文件做zip压缩，压缩文件做Base64编码后传输，字符编码UTF-8。
	ReceiptSumFile string `json:"receipt_sum_file"`
	// 电子回单信息集合
	ReceiptInfo []ReceiptInfo `json:"receiptInfo"`
}
