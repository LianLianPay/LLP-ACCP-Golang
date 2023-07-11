package refund

type RefundMethod struct {
	Method string  `json:"method"`
	Amount float64 `json:"amount"`
}
