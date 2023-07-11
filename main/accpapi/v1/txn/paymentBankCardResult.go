package txn

/**
 * 银行卡快捷支付 响应参数
 */
type PaymentBankCardResult struct {
	LinkedAgrtno   string `json:"linked_agrtno"`
	AccountingDate string `json:"accounting_date"`
	FinishTime     string `json:"finish_time"`
}
