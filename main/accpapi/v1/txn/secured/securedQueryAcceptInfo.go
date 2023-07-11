package secured

type SecuredQueryAcceptInfo struct {
	AccpTxNo string `json:"accp_txno"`
	TxnSeqNo string `json:"txn_seqno"`
	/*
	   担保确认单状态
	   PAYBILL_FINISH：支付完成，交易成功结束
	   WAIT_PAY：交易创建，等待付款
	   PRE_FINISH：预付完成
	   REFUND_MONEY：撤销
	   PAY_CLOSE： 支付关闭
	*/
	State string `json:"state"`
}
