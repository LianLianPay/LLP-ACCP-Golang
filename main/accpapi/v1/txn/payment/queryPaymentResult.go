package payment

/**
 * 支付结果查询 响应参数
 */
type QueryPaymentResult struct {
	RetCode    string `json:"ret_code"`
	RetMsg     string `json:"ret_msg"`
	OidPartner string `json:"oid_partner"`
	/*
	   交易类型。
	   用户充值：USER_TOPUP
	   商户充值：MCH_TOPUP
	   普通消费：GENERAL_CONSUME
	   担保消费：SECURED_CONSUME
	   担保确认：SECURED_CONFIRM
	   内部代发：INNER_FUND_EXCHANGE
	   外部代发：OUTER_FUND_EXCHANGE
	*/
	TxnType string `json:"txn_type"`
	// 账务日期
	AccountingDate string `json:"accounting_date"`
	// 支付完成时间
	FinishTime string `json:"finish_time"`
	// ACCP系统交易单号
	AccpTxNo float64 `json:"accp_txno"`
	// 渠道交易单号
	ChnlTxNo string `json:"chnl_txno"`
	/*
	   支付交易状态。
	   TRADE_WAIT_PAY:交易处理中
	   TRADE_SUCCESS:交易成功
	   TRADE_CLOSE:交易失败
	   支付交易状态以此为准，商户必须依据此字段值进行后续业务逻辑处理。
	*/
	TxnStatus string `json:"txn_status"`
	// 商户订单信息
	OrderInfo QueryPaymentOrderInfo `json:"orderInfo"`
	// 付款方信息（组合支付场景返回付款方信息数组）
	PayerInfo []QueryPaymentPayerInfo `json:"payerInfo"`
	// 收款方信息（交易分账场景返回收款方信息数组）
	PayeeInfo []QueryPaymentPayeeInfo `json:"payeeInfo"`
}
