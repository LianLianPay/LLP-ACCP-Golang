package txn

type PaymentBankCardInfo struct {
	LinkedAgrtno   string `json:"linked_agrtno"`
	LinkedAcctno   string `json:"linked_acctno"`
	LinkedPhone    string `json:"linked_phone"`
	LinkedAcctname string `json:"linked_acctname"`
	IDType         string `json:"id_type"`
	IDNo           string `json:"id_no"`
}
