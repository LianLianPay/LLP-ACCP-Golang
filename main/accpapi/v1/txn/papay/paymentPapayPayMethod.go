package papay

type PaymentPapayPayMethod struct {
	Method string  `json:"method"`
	Amount float64 `json:"amount"`
}
