package webapge

type OpenacctApplyLinkedAcctInfo struct {
	LinkedAcctType   string `json:"linked_accttype"`
	LinkedAcctNo     string `json:"linked_acctno"`
	LinkedBankCode   string `json:"linked_bankcode"`
	LinkedBrBankNo   string `json:"linked_brbankno"`
	LinkedBrBankName string `json:"linked_brbankname"`
	LinkedAcctName   string `json:"linked_acctname"`
	LinkedPhone      string `json:"linked_phone"`
}
