package secured

type CancelSecuredPaymentPayeeInfo struct {
	PayeeType    string  `json:"payee_type"`
	PayeeID      string  `json:"payee_id"`
	Amount       float64 `json:"amount"`
	AcceptAmount float64 `json:"accept_amount"`
	RefundAmount float64 `json:"refund_amount"`
}
