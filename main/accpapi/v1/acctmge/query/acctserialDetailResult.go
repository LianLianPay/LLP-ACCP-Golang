package query

/**
 * 资金流水详情查询 响应参数
 */
type AcctserialDetailResult struct {
	RetCode    string `json:"ret_code"`
	RetMsg     string `json:"ret_msg"`
	OidPartner string `json:"oid_partner"`
	DateAcct   string `json:"date_acct"`
	OidAcctno  string `json:"oid_acctno"`
	JnoAcct    string `json:"jno_acct"`
	TxnSeqno   string `json:"txn_seqno"`
	AccpTxnno  string `json:"accp_txnno"`
	/*
	   交易类型。
	   用户充值：USER_TOPUP
	   商户充值：MCH_TOPUP
	   普通消费：GENERAL_CONSUME
	   担保消费：SECURED_CONSUME
	   手续费收取：SERVICE_FEE
	   内部代发：INNER_FUND_EXCHANGE
	   外部代发：OUTER_FUND_EXCHANGE
	   账户提现：ACCT_CASH_OUT
	   担保确认：SECURED_CONFIRM
	   手续费应收应付核销：CAPITAL_CANCEL
	   定向内部代发：INNER_DIRECT_EXCHANGE
	*/
	TxnType       string `json:"txn_type"`
	ProductCode   string `json:"product_code"`
	OtherAcct     string `json:"other_acct"`
	OtherAcctName string `json:"other_acct_name"`
	TxnTime       string `json:"txn_time"`
	FlagDC        string `json:"flag_dc"`
	Amount        string `json:"amt"`
	AmountBal     string `json:"amt_bal"`
	Memo          string `json:"memo"`
	JnoCli        string `json:"jno_cli"`
}
