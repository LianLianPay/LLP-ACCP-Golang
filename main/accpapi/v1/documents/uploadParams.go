package documents

/**
 * 文件上传 请求参数
 */
type UploadParams struct {
	Timestamp  string `json:"timestamp"`
	OidPartner string `json:"oid_partner"`
	UserID     string `json:"user_id"`
	TxnSeqno   string `json:"txn_seqno"`
	TxnTime    string `json:"txn_time"`
	// 文件类型。支持bmp、png、jpeg、jpg、gif
	FileType    string `json:"file_type"`
	ContextType string `json:"context_type"`
	/*
	   文件名称。内容类型为SUPPLEMENT_CSV时必填
	   格式为SUPPLEMENT_CSV_商户号_对账日期.csv注：文件名请使用大写字母
	*/
	FileName string `json:"file_name"`
	/*
	   文件内容。文件流通过Base64编码后传输，字符编码UTF-8；最大6M。
	   Base64编码后的字符串需去掉前缀，前缀示例：data:image/png;base64
	*/
	FileContext string `json:"file_context"`
}
