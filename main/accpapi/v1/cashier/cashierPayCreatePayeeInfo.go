package cashier

type CashierPayCreatePayeeInfo struct {
	PayeeID       string `json:"payee_id"`
	PayeeType     string `json:"payee_type"`
	PayeeAcctType string `json:"payee_accttype"`
	PayeeAmount   string `json:"payee_amount"`
	PayeeMemo     string `json:"payee_memo"`
}
