package txn

type PaymentBankCardPayMethods struct {
	Method string  `json:"method"`
	Amount float64 `json:"amount"`
}
