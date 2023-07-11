package transfer

/**
 * 内部代发申请 响应参数
 */
type TransferMorepyeeResult struct {
	RetCode    string `json:"ret_code"`
	RetMsg     string `json:"ret_msg"`
	OidPartner string `json:"oid_partner"`
	UserID     string `json:"user_id"`
	TxnSeqNo   string `json:"txn_seqno"`
	// 订单总金额，单位为元，精确到小数点后两位
	TotalAmount float64 `json:"total_amount"`
	AccpTxNo    string  `json:"accp_txno"`
	// 支付授权令牌，有效期30分钟。当交易需要二次验证的时候，需要通过token调用调用交易二次短信验证接口
	Token string `json:"token"`
	// 账务日期。ACCP系统交易账务日期，交易成功时返回，格式：yyyyMMdd
	AccountingDate string `json:"accounting_date"`
}
