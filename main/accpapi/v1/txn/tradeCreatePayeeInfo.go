package txn

type TradeCreatePayeeInfo struct {
	PayeeID       string `json:"payee_id"`
	PayeeType     string `json:"payee_type"`
	PayeeAcctType string `json:"payee_acct_type"`
	PayeeAmount   string `json:"payee_amount"`
	PayeeMemo     string `json:"payee_memo"`
}
